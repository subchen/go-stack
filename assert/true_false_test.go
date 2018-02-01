package assert

import (
	"testing"
)

func TestTrue(t *testing.T) {
	True(t, true)
}

func TestFalse(t *testing.T) {
	False(t, false)
}
