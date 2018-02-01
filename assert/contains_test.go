package assert

import (
	"testing"
)

func TestContains(t *testing.T) {
	Contains(t, "abc", "a")
	Contains(t, []string{"a", "b"}, "a")
	Contains(t, map[string]string{"a": "1"}, "a")
}

func TestNotContains(t *testing.T) {
	NotContains(t, "abc", "x")
	NotContains(t, []string{"a", "b"}, "x")
	NotContains(t, map[string]string{"a": "1"}, "x")
}
