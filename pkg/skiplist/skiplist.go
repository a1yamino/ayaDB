package skiplist

type SkipList struct {
	header *node
	tail   *node
	level  int
	length int
}

type node struct{}
