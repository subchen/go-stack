package assert

import (
	"testing"
)

func TestZero(t *testing.T) {
	Zero(t, nil)
	Zero(t, "")
	Zero(t, 0)
	Zero(t, uint64(0))
	Zero(t, false)
}

func TestNotZero(t *testing.T) {
	NotZero(t, "123")
	NotZero(t, 1)
	NotZero(t, true)
	NotZero(t, make([]string, 0))
}
