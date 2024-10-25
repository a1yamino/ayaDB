package ayadb

type Iterater interface {
	Next()
	Valid() bool
	Rewind()
	Item() *Item
	Close() error
	// Seek(key string)
}

type Item struct {
	key   []byte
	value []byte
}

func (i Item) Key() []byte {
	return i.key
}
func (i Item) Value() []byte {
	return i.value
}

type Options struct {
	Prefix []byte
	IsAsc  bool
}

func NewIterater(opt *Options) Iterater {
	return nil
}

type IteraterImpl struct {
}

func (i *IteraterImpl) Next() {
}

func (i *IteraterImpl) Valid() bool {
	return false
}

func (i *IteraterImpl) Rewind() {
}

func (i *IteraterImpl) Item() *Item {
	return nil
}

func (i *IteraterImpl) Close() error {
	return nil
}
