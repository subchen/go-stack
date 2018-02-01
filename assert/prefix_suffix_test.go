package assert

import (
	"testing"
)

func TestHasPrefix(t *testing.T) {
	HasPrefix(t, "", "")
	HasPrefix(t, "abc", "")
	HasPrefix(t, "abc", "a")
	HasPrefix(t, "abc", "ab")
}

func TestHasNotPrefix(t *testing.T) {
	HasNotPrefix(t, "", "b")
	HasNotPrefix(t, "abc", "b")
	HasNotPrefix(t, "abc", "abcd")
	HasNotPrefix(t, "abc", "dd")
}

func TestHasSuffix(t *testing.T) {
	HasSuffix(t, "", "")
	HasSuffix(t, "abc", "")
	HasSuffix(t, "abc", "c")
	HasSuffix(t, "abc", "bc")
}

func TestHasNotSuffix(t *testing.T) {
	HasNotSuffix(t, "", "b")
	HasNotSuffix(t, "abc", "b")
	HasNotSuffix(t, "abc", "0abc")
	HasNotSuffix(t, "abc", "dd")
}
