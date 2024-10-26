package codec

import "time"

type Entry struct {
	Key      []byte
	Value    []byte
	ExpireAt uint64
}

// NewEntry
func NewEntry(key, value []byte) *Entry {
	return &Entry{
		Key:   key,
		Value: value,
	}
}

func (e *Entry) WithTTL(duration time.Duration) *Entry {
	e.ExpireAt = uint64(time.Now().Add(duration).Unix())
	return e
}
