package fs

import (
	"testing"
)

func TestIsSymlink(t *testing.T) {
	var tests = []struct {
		path string
		want bool
	}{
		{"/proc/self/exe", true},
		{"/home", false},
	}

	for _, test := range tests {
		if got := IsSymlink(test.path); got != test.want {
			t.Errorf("IsSymlink(%q) = %v; want: %v", test.path, got, test.want)
		}
	}
}
