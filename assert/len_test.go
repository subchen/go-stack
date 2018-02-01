package assert

import "testing"

func TestLen(t *testing.T) {
	Len(t, "", 0)
	Len(t, make([]string, 10), 10)
	Len(t, make(map[string]string), 0)
	Len(t, make(chan int), 0)
}
