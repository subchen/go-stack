package assert

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	Empty(t, nil)
	Empty(t, "")
	Empty(t, 0)
	Empty(t, uint64(0))
	Empty(t, float64(0))
	Empty(t, []string{})
	Empty(t, map[string]string{})
	Empty(t, make([]int, 0))
	Empty(t, make(map[int]int))
	Empty(t, make(chan int))
}

func TestNotEmpty(t *testing.T) {
	NotEmpty(t, "123")
	NotEmpty(t, 123)
	NotEmpty(t, []string{"123"})
	NotEmpty(t, map[string]string{"abc": "123"})
	NotEmpty(t, make([]int, 10))
}
