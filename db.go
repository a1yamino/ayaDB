package ayadb

import (
	"ayaDB/pkg/file"
	"ayaDB/pkg/lsm"
)

type KVDB interface {
	Set(data *Entry) error
	Get(key []byte) (*Entry, error)
	Delete(key []byte) error
	NewIterater(opt *Options) Iterater
	Info() *Stats
	Close() error
}

var _ KVDB = &DB{}

type Entry struct {
	Key   []byte
	Value []byte
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
	// memtable flush and clear wal log
	return nil
}

func (db *DB) Set(data *Entry) error {
	// write to lsm
	return nil
}

func (db *DB) Get(key []byte) (*Entry, error) {
	// read from memtable
	// if not found, read from lsm
	return nil, nil
}

func (db *DB) Delete(key []byte) error {
	// write a null value to lsm as tombstone
	return nil
}

func (db *DB) NewIterater(opt *Options) Iterater {
	return NewIterater(opt)
}

func (db *DB) Info() *Stats {
	return db.stats
}
