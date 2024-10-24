package ayadb

import "testing"

func TestKVDB(t *testing.T) {
	db := NewDB()
	key := "test"
	value := "test_value"

	err := db.Set(key, value)
	if err != nil {
		t.Errorf("Set failed: %v", err)
	}

	v, err := db.Get(key)
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if v != value {
		t.Errorf("Get returned %s, want %s", v, value)
	}

	err = db.Delete(key)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}

	v, err = db.Get(key)
	if err == nil {
		t.Errorf("Get succeeded, want error")
	}
	if v != "" {
		t.Errorf("Get returned %s, want empty string", v)
	}
}
