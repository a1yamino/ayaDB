package ayadb

import (
	"ayaDB/pkg/codec"
	"ayaDB/pkg/file"
	"ayaDB/pkg/iterater"
	"ayaDB/pkg/lsm"
	"ayaDB/pkg/utils"
)

type KVDB interface {
	Set(data *codec.Entry) error
	Get(key []byte) (*codec.Entry, error)
	Delete(key []byte) error
	// NewIterater(opt *iterater.Options) iterater.Iterater
	Info() *Stats
	Close() error
}

var _ KVDB = &DB{}

type Options struct {
	ValueThreshold int64
}

type DB struct {
	opt   *Options
	lsm   *lsm.LSM
	vlog  *file.VLog
	stats *Stats
}

func Open(opt *Options) *DB {
	db := &DB{opt: opt}
	db.lsm = lsm.NewLSM(&lsm.Options{})
	db.vlog = file.NewVLog(&file.Options{})
	db.stats = newStats(opt)
	go db.lsm.StartMerge()
	go db.vlog.StartGC()
	go db.stats.StartStats()
	return db
}

func (db *DB) Close() error {
	if err := db.lsm.Close(); err != nil {
		return err
	}
	if err := db.vlog.Close(); err != nil {
		return err
	}
	return db.stats.close()
}

func (db *DB) Set(data *codec.Entry) error {
	// if value size > threshold, write to vlog
	vp := &codec.ValuePtr{}
	if utils.ValueSize(data.Value) > db.opt.ValueThreshold {
		vp = codec.NewValuePtr(data)
		// no problem if vlog write failed
		// if failed to write to lsm, we can clear invalid vlog with gc
		if err := db.vlog.Write(data); err != nil {
			return err
		}
	}
	// write to lsm
	if vp != nil {
		data.Value = codec.ValuePtrCodec(vp)
	}
	return db.lsm.Set(data)
}

func (db *DB) Get(key []byte) (entry *codec.Entry, err error) {
	// read from memtable
	// if not found, read from lsm
	if entry, err = db.lsm.Get(key); err == nil {
		return
	}
	// read from vlog
	if entry != nil && codec.IsValuePtr(entry) {
		if entry, err = db.vlog.Read(); err == nil {
			return
		}
	}
	return nil, nil
}

func (db *DB) Delete(key []byte) error {
	// write a null value to lsm as tombstone
	db.Set(&codec.Entry{Key: key, Value: nil})
	return nil
}

func (db *DB) NewIterater(opt *iterater.Options) iterater.Iterater {
	return iterater.NewIterater(opt)
}

func (db *DB) Info() *Stats {
	return db.stats
}
