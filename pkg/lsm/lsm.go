package lsm

type LSM struct {
	memTable   *memTable
	immutables []*memTable
	levels     *levelManager
	opt        *Options
}
type Options struct {
}

func NewLSM(opt *Options) *LSM {
	lsm := &LSM{opt: opt}
	lsm.memTable, lsm.immutables = recovery(opt)
	lsm.levels = newLevelManager(opt)
	return lsm
}

func (lsm *LSM) StartMerge() {}
