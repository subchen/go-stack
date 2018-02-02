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
		want   []string
	}{
		{"", ".", "[]", nil},
		{"a.b.c", ".", "[]", []string{"a", "b", "c"}},
		{"a.[b].c", ".", "[]", []string{"a", "b", "c"}},
		{"a.[b.c]", ".", "[]", []string{"a", "b.c"}},
		{"a.[b.c].dd  ", ".", "[]", []string{"a", "b.c", "dd"}},
		{` select * from   "my table" `, " ", `"`, []string{"select", "*", "from", "my table"}},
	}

	for _, test := range tests {
		if got := SplitWithQuotes(test.str, test.sep, test.quotes); !reflect.DeepEqual(got, test.want) {
			t.Errorf("SplitWithQuotes(%q, %q, %q) = %q; want: %q", test.str, test.sep, test.quotes, got, test.want)
		}
	}
}

