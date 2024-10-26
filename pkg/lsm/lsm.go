package lsm

import (
	"ayaDB/pkg/codec"
	"ayaDB/pkg/utils"
)

type LSM struct {
	memTable   *memTable
	immutables []*memTable
	levels     *levelManager
	opt        *Options
	closer     *utils.Closer
}
type Options struct {
}

func NewLSM(opt *Options) *LSM {
	lsm := &LSM{opt: opt}
	lsm.memTable, lsm.immutables = recovery(opt)
	lsm.levels = newLevelManager(opt)
	return lsm
}

func (lsm *LSM) Close() error {
	if err := lsm.memTable.close(); err != nil {
		return err
	}
	for _, im := range lsm.immutables {
		if err := im.close(); err != nil {
			return err
		}
	}
	if err := lsm.levels.close(); err != nil {
		return err
	}
	lsm.closer.Close()
	return nil
}

func (lsm *LSM) StartMerge() {
	defer lsm.closer.Done()
	for {
		select {
		case <-lsm.closer.Wait():
			return
		default:
			// merge
		}
	}
}

func (lsm *LSM) Set(entry *codec.Entry) error {
	// check current memtable if full, flush to immutable
	// write to memtable
	if err := lsm.memTable.set(entry); err != nil {
		return err
	}
	// check if exists any immutable need to be flushed
	for _, im := range lsm.immutables {
		if err := lsm.levels.flush(im); err != nil {
			return err
		}
	}
	return nil
}

func (lsm *LSM) Get(key []byte) (*codec.Entry, error) {
	// read from memtable
	if entry, err := lsm.memTable.get(key); err == nil {
		return entry, nil
	}
	// read from immutable
	for _, im := range lsm.immutables {
		if entry, err := im.get(key); err == nil {
			return entry, nil
		}
	}
	// read from levels
	return lsm.levels.Get(key)
}
