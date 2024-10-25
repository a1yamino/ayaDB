package lsm

import (
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
