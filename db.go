package ayadb

import (
	"fmt"
	"sync"
)

type KVDB interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

var _ KVDB = &DB{}

type DB struct {
	sync.RWMutex
	storage map[string]string
}

func NewDB() *DB {
	return &DB{storage: make(map[string]string)}
}

func (db *DB) Set(key string, value string) error {
	db.storage[key] = value
	return nil
}

func (db *DB) Get(key string) (string, error) {
	value, ok := db.storage[key]
	if !ok {
		return "", fmt.Errorf("key %s not found", key)
	}
	return value, nil
}

func (db *DB) Delete(key string) error {
	delete(db.storage, key)
	return nil
}
