package strmap

import "testing"

func TestCopyOnWriteMap(t *testing.T) {
	m := NewCopyOnWriteMap()

	m.Copy(map[Key]Value{
		"a": "1",
		"b": "2",
		"c": "3",
	})

	m.Set("d", "4")

	if m.Get("d") != "4" {
		t.Fail()
	}

	if _, ok := m.GetOK("e"); ok {
		t.Fail()
	}

	m.Remove("d")

	if _, ok := m.GetOK("d"); ok {
		t.Fail()
	}

	m.Clear()

	if _, ok := m.GetOK("a"); ok {
		t.Fail()
	}

	m.Set("e", "5")

	if m.Get("e") != "5" {
		t.Fail()
	}
}
