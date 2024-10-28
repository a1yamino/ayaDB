package lsm

import "testing"

func TestLevels(t *testing.T) {
	levels := newLevelManager(&Options{})
	defer levels.close()
	if entry, err := levels.Get([]byte("test")); err != nil {
		t.Error(err)
	} else {
		t.Logf("levels.Get: %v", entry)
	}
}
