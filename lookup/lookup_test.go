package lookup

import "testing"

func TestFind(t *testing.T) {
	filename := "lookup.go"
	if _, err := Find(filename); err != nil {
		t.Errorf("Find(%v) not found", filename)
	}
}

func TestFindGlob(t *testing.T) {
	filename := "lookup.*"
	if _, err := Find(filename); err != nil {
		t.Errorf("Find(%v) not found", filename)
	}
}

func TestFind_notFound(t *testing.T) {
	filename := "not.found"
	if _, err := Find(filename); err == nil {
		t.Errorf("Find(%v) was found", filename)
	}
}
