package assert

import (
	"testing"
)

func TestPanic(t *testing.T) {
	Panic(t, func() {
		panic("error")
	})
}

func TestNoPanic(t *testing.T) {
	NoPanic(t, func() {
		//do nothing
	})
}
