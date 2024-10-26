package ayadb

import "ayaDB/pkg/utils"

type Stats struct {
	closer     *utils.Closer
	EntryCount int64 // count of entries
}

// Close
func (s *Stats) close() error {
	return nil
}

// StartStats
func (s *Stats) StartStats() {
	defer s.closer.Done()
	for {
		select {
		case <-s.closer.Wait():
			return
		default:
			// stats
		}
	}
}

// NewStats
func newStats(opt *Options) *Stats {
	s := &Stats{}
	s.closer = utils.NewCloser(1)
	// temporary set EntryCount to 1
	s.EntryCount = 1
	return s
}
