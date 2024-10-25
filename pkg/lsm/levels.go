package lsm

import "ayaDB/pkg/file"

type levelManager struct {
	opt      *Options
	cache    *cache
	manifest *file.Manifest
	levels   []*levelHandler
}

type levelHandler struct {
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

func (lm *levelManager) loadManifest() {
	lm.manifest = file.OpenManifestFile(&file.Options{})
}

func (lm *levelManager) build() {
	lm.levels = make([]*levelHandler, 0)
	lm.levels = append(lm.levels, &levelHandler{tables: []*table{openTable(lm.opt)}})
	lm.loadCache()
}
