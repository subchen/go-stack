package maps

import "testing"

func TestStringInterfaceCOWMap(t *testing.T) {
	m := NewStringInterfaceCOWMap()

	m.CopyFrom(map[string]interface{}{
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
