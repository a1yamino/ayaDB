package ayadb

type Iterater interface {
	Next()
	Valid() bool
	Rewind()
	Item() Item
	Close() error
	Seek(key string)
}

type Item struct {
	key   string
	value string
}
