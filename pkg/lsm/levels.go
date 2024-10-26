package lsm

import (
	"ayaDB/pkg/codec"
	"ayaDB/pkg/file"
	"ayaDB/pkg/utils"
)

type levelManager struct {
	opt      *Options
	cache    *cache
	manifest *file.Manifest
	levels   []*levelHandler
}

type levelHandler struct {
	level  int
	tables []*table
}

func newLevelManager(opt *Options) *levelManager {
	lm := &levelManager{opt: opt}
	lm.loadManifest()
	lm.build()
	return lm
}

func (lm *levelManager) loadCache() {
}

func (lm *levelManager) close() error {
	if err := lm.cache.close(); err != nil {
		return err
	}
	if err := lm.manifest.Close(); err != nil {
		return err
	}
	for _, lh := range lm.levels {
		if err := lh.close(); err != nil {
			return err
		}
	}
	return nil
}

func (lm *levelManager) loadManifest() {
	lm.manifest = file.OpenManifestFile(&file.Options{})
}

func (lm *levelManager) build() {
	lm.levels = make([]*levelHandler, 8)
	lm.levels[0] = &levelHandler{tables: []*table{openTable(lm.opt)}, level: 0}
	for i := 1; i < utils.MaxLevelNum; i++ {
		lm.levels[i] = &levelHandler{level: i}
	}
	lm.loadCache()
}

func (lm *levelManager) flush(immutable *memTable) error {
	// flush immutable to level 0
	return nil
}

func (lm *levelManager) Get(key []byte) (entry *codec.Entry, err error) {
	if entry, err = lm.levels[0].Get(key); err == nil {
		return
	}
	for i := 1; i < utils.MaxLevelNum; i++ {
		ld := lm.levels[i]
		if entry, err = ld.Get(key); err == nil {
			return
		}
	}
}

func (lh *levelHandler) Get(key []byte) (*codec.Entry, error) {
	return nil, nil
}

func (lh *levelHandler) close() error {
	return nil
}
