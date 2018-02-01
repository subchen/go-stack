package maps

import "testing"

func TestInterfaceInterfaceMap(t *testing.T) {
	m := NewInterfaceInterfaceMap()

	m.CopyFrom(map[interface{}]interface{}{
		"a": "1",
		"b": "2",
		"c": "3",
	})

	m.Set("d", "4")

	if m.Get("d") != "4" {
		t.Fail()
	}

	if _, ok := m.Load("e"); ok {
		t.Fail()
	}

	m.Remove("d")

	if _, ok := m.Load("d"); ok {
		t.Fail()
	}

	m.Clear()

	if _, ok := m.Load("a"); ok {
		t.Fail()
	}

	m.Set("e", "5")

	if m.Get("e") != "5" {
		t.Fail()
	}
}
