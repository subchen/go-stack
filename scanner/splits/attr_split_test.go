package splits

import (
	"reflect"
	"testing"
)

func TestAttrSplit(t *testing.T) {
	input := `nodes."availables".[0].status.[cpu=2.0]`
	attrs, err := AttrSplit(input)
	if err != nil {
		t.Errorf("Got err: %v", err)
		return
	}

	want := []string{
		"nodes",
		"availables",
		"[0]",
		"status",
		"[cpu=2.0]",
	}
	if !reflect.DeepEqual(attrs, want) {
		t.Errorf("splits error, got: %v", attrs)
		return
	}
}
