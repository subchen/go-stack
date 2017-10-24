package gstack

import "testing"

func TestPathGetNameWithoutExt(t *testing.T) {
	var tests = []struct {
		path string
		want string
	}{
		{"abc.txt", "abc"},
		{"/etc/abc.txt", "abc"},
		{"/etc/abc.txt.tmpl", "abc.txt"},
	}

	for _, test := range tests {
		if got := PathGetNameWithoutExt(test.path); got != test.want {
			t.Errorf("PathGetNameWithoutExt(%q) = %q; want: %q", test.path, got, test.want)
		}
	}
}
