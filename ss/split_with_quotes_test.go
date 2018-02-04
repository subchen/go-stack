package ss

import (
	"reflect"
	"testing"
)

func TestSplitWithQuotes(t *testing.T) {
	var tests = []struct {
		str    string
		sep    string
		quotes string
		keep   bool
		want   []string
	}{
		{"", ".", "[]", false, nil},
		{"a.b.c", ".", "[]", false, []string{"a", "b", "c"}},
		{"a.[b].c", ".", "[]", false, []string{"a", "b", "c"}},
		{"a.[b].c", ".", "[]", true, []string{"a", "[b]", "c"}},
		{"a.[b.c]", ".", "[]", false, []string{"a", "b.c"}},
		{"a.[b.c]", ".", "[]", true, []string{"a", "[b.c]"}},
		{"a.[b.c].dd  ", ".", "[]", false, []string{"a", "b.c", "dd"}},
		{` select * from   "my table" `, " ", `"`, false, []string{"select", "*", "from", "my table"}},
	}

	for _, test := range tests {
		if got := SplitWithQuotes(test.str, test.sep, test.quotes, test.keep); !reflect.DeepEqual(got, test.want) {
			t.Errorf("SplitWithQuotes(%q, %q, %q, %q) = %q; want: %q", test.str, test.sep, test.quotes, test.keep, got, test.want)
		}
	}
}
