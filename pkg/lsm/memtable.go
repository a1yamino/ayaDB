package lsm

import (
	"ayaDB/pkg/codec"
	"ayaDB/pkg/file"
	"ayaDB/pkg/skiplist"
)

type memTable struct {
	wal *file.WalFile
	sl  *skiplist.SkipList
}

// recovery
func recovery(opt *Options) (*memTable, []*memTable) {
	fileOpt := &file.Options{}
	return &memTable{wal: file.OpenWalFile(fileOpt)}, []*memTable{}
}

func (m *memTable) close() error {
	if err := m.wal.Close(); err != nil {
		return err
	}
	if err := m.sl.Close(); err != nil {
		return err
	}
	return nil
}

func (m *memTable) set(entry *codec.Entry) error {
	// write to wal
	if err := m.wal.Write(entry); err != nil {
		return err
	}
	// write to skiplist
	return m.sl.Insert(entry)
}

func (m *memTable) get(key []byte) (*codec.Entry, error) {
	return m.sl.Search(key), nil
}
