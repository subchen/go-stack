package fs

import (
	"testing"
)

func TestBasenameWithoutExt(t *testing.T) {
	var tests = []struct {
		path string
		want string
	}{
		{"abc.txt", "abc"},
		{"/etc/abc.txt", "abc"},
		{"/etc/abc.txt.tmpl", "abc.txt"},
	}

	for _, test := range tests {
		if got := BasenameWithoutExt(test.path); got != test.want {
			t.Errorf("BasenameWithoutExt(%q) = %q; want: %q", test.path, got, test.want)
		}
	}
}
