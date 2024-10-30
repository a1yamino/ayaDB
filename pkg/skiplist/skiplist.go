package skiplist

import (
	"ayaDB/pkg/codec"
	"ayaDB/pkg/utils"
	"bytes"
	"math/rand"
	"time"
)

type SkipList struct {
	header *node
	rand   *rand.Rand
	level  int
	length int
}

func (sl *SkipList) Close() error {
	return nil
}

type node struct {
	data   *codec.Entry
	levels []*node
	score  float64
}

func newNode(entry *codec.Entry, level int, score float64) *node {
	return &node{
		data:   entry,
		levels: make([]*node, level),
		score:  score,
	}
}

func (sl *SkipList) Insert(entry *codec.Entry) error {
	score := sl.calculateScore(entry.Key)
	var nd *node

	max := len(sl.header.levels)
	prev := sl.header

	var prevHeads [utils.DefaultMaxSkipListLevels]*node
	for i := max - 1; i >= 0; {
		prevHeads[i] = prev

		for next := prev.levels[i]; next != nil; next = prev.levels[i] {
			if comp := sl.compare(score, entry.Key, next); comp > 0 {
				prev = next
				prevHeads[i] = prev
			} else {
				if comp == 0 {
					nd = next
					nd.data.Value = entry.Value
					return nil
				}
				break
			}

		}

		topLevel := prev.levels[i]

		for i--; i >= 0 && prev.levels[i] == topLevel; i-- {
			prevHeads[i] = prev
		}
	}
	level := sl.randomLevel()

	nd = newNode(entry, level, score)

	for i := 0; i < level; i++ {
		nd.levels[i] = prevHeads[i].levels[i]
		prevHeads[i].levels[i] = nd
	}
	sl.length++
	return nil
}

func (sl *SkipList) Search(key []byte) *codec.Entry {
	if sl.length == 0 {
		return nil
	}
	score := sl.calculateScore(key)

	prev := sl.header
	i := len(sl.header.levels) - 1
	for i >= 0 {
		for next := prev.levels[i]; next != nil; next = prev.levels[i] {
			if comp := sl.compare(score, key, next); comp > 0 {
				prev = next
			} else if comp == 0 {
				return next.data
			} else {
				break
			}
		}

		topLevel := prev.levels[i]
		for i--; i >= 0 && prev.levels[i] == topLevel; i-- {
		}
	}
	return nil
}

func (sl *SkipList) Delete(key []byte) error {
	score := sl.calculateScore(key)

	max := len(sl.header.levels)
	prev := sl.header

	var (
		prevHeads [utils.DefaultMaxSkipListLevels]*node
		nd        *node
	)

	for i := max - 1; i >= 0; {
		prevHeads[i] = prev

		for next := prev.levels[i]; next != nil; next = prev.levels[i] {
			if comp := sl.compare(score, key, next); comp > 0 {
				prev = next
				prevHeads[i] = prev
			} else {
				if comp == 0 {
					nd = next
					break
				}
				break
			}
		}

		topLevel := prev.levels[i]

		for i--; i >= 0 && prev.levels[i] == topLevel; i-- {
			prevHeads[i] = prev
		}
	}
	if nd == nil {
		return nil
	}

	prevTopLevel := len(prev.levels)
	for i := 0; i < prevTopLevel; i++ {
		prevHeads[i].levels[i] = nd.levels[i]
	}

	sl.length--
	return nil
}

func NewSkipList() *SkipList {
	source := rand.NewSource(time.Now().UnixNano())
	return &SkipList{
		header: &node{
			levels: make([]*node, utils.DefaultMaxSkipListLevels),
		},
		rand:   rand.New(source),
		level:  utils.DefaultMaxSkipListLevels,
		length: 0,
	}
}

func (sl *SkipList) calculateScore(key []byte) float64 {
	var hash uint64
	l := len(key)
	if l > 8 {
		l = 8
	}
	for i := 0; i < l; i++ {
		shift := uint(64 - 8*(i+1))
		hash |= uint64(key[i]) << shift
	}

	return float64(hash)
}

// compare returns 1 if score is greater than next.score, -1 if score is less than next.score, 0 if equal
func (sl *SkipList) compare(score float64, key []byte, next *node) int {
	if score > next.score {
		return 1
	} else if score < next.score {
		return -1
	} else {
		return bytes.Compare(key, next.data.Key)
	}
}

func (sl *SkipList) randomLevel() int {
	level := 1
	for level < sl.level && sl.rand.Intn(2) == 0 {
		level++
	}
	return level
}
