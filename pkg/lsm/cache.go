package lsm

type cache struct {
}

// NewCache
func newCache(opt *Options) *cache {
	return &cache{}
}

func (c *cache) close() error {
	return nil
}
