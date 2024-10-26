package skiplist

import (
	"ayaDB/pkg/codec"
)

type SkipList struct {
	header *node
	tail   *node
	level  int
	length int
}

func (sl *SkipList) Close() error {
	return nil
}

type node struct {
	data   *codec.Entry
	prev   *node
	next   *node
	levels []*node
}

func (sl *SkipList) Insert(entry *codec.Entry) error {
	sl.header.next = &node{data: entry}
	return nil
}

func (sl *SkipList) Search(key []byte) *codec.Entry {
	return sl.header.data
}

func NewSkipList() *SkipList {
	return &SkipList{
		header: &node{},
	}
}
