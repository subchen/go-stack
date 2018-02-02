package ss

import "testing"

func TestStringBefore(t *testing.T) {
	var tests = []struct {
		str  string
		find string
		want string
	}{
		{"", "abc", ""},
		{"abc", "a", ""},
		{"abc", "b", "a"},
		{"abc", "c", "ab"},
		{"abc", "d", "abc"},
		{"abc", "", ""},
	}

	for _, test := range tests {
		if got := StringBefore(test.str, test.find); got != test.want {
			t.Errorf("StringBefore(%q, %q) = %q; want: %q", test.str, test.find, got, test.want)
		}
	}
}

func TestStringAfter(t *testing.T) {
	var tests = []struct {
		str  string
		find string
		want string
	}{
		{"", "*", ""},
		{"abc", "a", "bc"},
		{"abcba", "b", "cba"},
		{"abc", "c", ""},
		{"abc", "d", ""},
		{"abc", "", "abc"},
	}

	for _, test := range tests {
		if got := StringAfter(test.str, test.find); got != test.want {
			t.Errorf("StringAfter(%q, %q) = %q; want: %q", test.str, test.find, got, test.want)
		}
	}
}

func TestStringBeforeLast(t *testing.T) {
	var tests = []struct {
		str  string
		find string
		want string
	}{
		{"", "*", ""},
		{"abcba", "b", "abc"},
		{"abc", "c", "ab"},
		{"a", "a", ""},
		{"a", "z", "a"},
		{"a", "", "a"},
	}

	for _, test := range tests {
		if got := StringBeforeLast(test.str, test.find); got != test.want {
			t.Errorf("StringBeforeLast(%q, %q) = %q; want: %q", test.str, test.find, got, test.want)
		}
	}
}

func TestStringAfterLast(t *testing.T) {
	var tests = []struct {
		str  string
		find string
		want string
	}{
		{"", "*", ""},
		{"*", "", ""},
		{"abc", "a", "bc"},
		{"abcba", "b", "a"},
		{"abc", "c", ""},
		{"a", "a", ""},
		{"a", "z", ""},
	}

	for _, test := range tests {
		if got := StringAfterLast(test.str, test.find); got != test.want {
			t.Errorf("StringAfterLast(%q, %q) = %q; want: %q", test.str, test.find, got, test.want)
		}
	}
}

func TestStringBetweenLast(t *testing.T) {
	var tests = []struct {
		str   string
		start string
		end   string
		want  string
	}{
		{"", "*", "*", ""},
		{"*", "", "", ""},
		{"abc", "a", "c", "b"},
		{"abcba", "a", "a", "bcb"},
		{"abcba", "a", "b", ""},
		{"abc", "c", "a", ""},
		{"abc", "a", "d", ""},
		{"abc", "z", "a", ""},
	}

	for _, test := range tests {
		if got := StringBetween(test.str, test.start, test.end); got != test.want {
			t.Errorf("StringAfterLast(%q, %q, %q) = %q; want: %q", test.str, test.start, test.end, got, test.want)
		}
	}
}
