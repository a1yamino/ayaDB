package utils

import "sync"

type Map struct {
	m sync.Map
}

func NewMap() *Map {
	return &Map{m: sync.Map{}}
}

func (m *Map) Get(key interface{}) (interface{}, bool) {
	return m.m.Load(key)
}

func (m *Map) Set(key, value interface{}) {
	m.m.Store(key, value)
}

func (m *Map) Range(f func(key, value interface{}) bool) {
	m.m.Range(f)
}
