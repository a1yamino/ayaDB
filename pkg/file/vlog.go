package file

import (
	"ayaDB/pkg/codec"
	"ayaDB/pkg/utils"
)

// VLog
type VLog struct {
	closer *utils.Closer
}

// NewVLog
func NewVLog(opt *Options) *VLog {
	v := &VLog{}
	v.closer = utils.NewCloser(1)
	return v
}

// StartGC
func (v *VLog) StartGC() {
	defer v.closer.Done()
	for {
		select {
		case <-v.closer.Wait():
			return
		default:
			// gc
		}
	}
}

func (v *VLog) Close() error {
	v.closer.Close()
	return nil
}

func (v *VLog) Write(entry *codec.Entry) error {
	return nil
}

func (v *VLog) Read() (*codec.Entry, error) {
	return nil, nil
}
