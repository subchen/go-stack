package assert

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	Error(t, errors.New("test"))
}

func TestNotError(t *testing.T) {
	NoError(t, nil)
}
