package lsm

import "ayaDB/pkg/utils"

type cache struct {
	indexs utils.Map // key: fid, value: tablebuffer
	blocks utils.Map // key: cacheID_blockoffset, value: block
}

type tableBuffer struct {
	t       *table
	cacheID int64
}

type blockBuffer struct {
	b []byte
}

// NewCache
func newCache(opt *Options) *cache {
	return &cache{}
}

func (c *cache) close() error {
	return nil
}
