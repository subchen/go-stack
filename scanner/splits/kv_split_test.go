package splits

import (
	"testing"
)

func TestKVSplitUsingColon(t *testing.T) {
	input := `exe:/bin/ls args:"-l -a --color"`
	m, err := KVSplit(input)
	if err != nil {
		t.Errorf("Got err: %v", err)
		return
	}

	if v := m["exe"]; v != "/bin/ls" {
		t.Errorf("Value for exe, got: %v", v)
		return
	}

	if v := m["args"]; v != "-l -a --color" {
		t.Errorf("Value for args, got: %v", v)
		return
	}
}

func TestKVSplitUsingEqual(t *testing.T) {
	input := `exe=/bin/ls args="-l -a --color"`
	m, err := KVSplit(input)
	if err != nil {
		t.Errorf("Got err: %v", err)
		return
	}

	if v := m["exe"]; v != "/bin/ls" {
		t.Errorf("Value for exe, got: %v", v)
		return
	}

	if v := m["args"]; v != "-l -a --color" {
		t.Errorf("Value for args, got: %v", v)
		return
	}
}
