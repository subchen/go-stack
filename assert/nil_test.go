package assert

import (
	"testing"
)

func TestNil(t *testing.T) {
	Nil(t, nil)
}

func TestNotNil(t *testing.T) {
	NotNil(t, "")
	NotNil(t, false)
	NotNil(t, []string{})
}
