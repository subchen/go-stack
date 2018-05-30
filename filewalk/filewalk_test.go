package filewalk

import (
	"os"
	"strings"
	"testing"
)

func TestFindFiles(t *testing.T) {
	acceptFn := func(path string, info os.FileInfo) bool {
		return strings.HasSuffix(info.Name(), ".go")
	}
	skipDirFn := func(path string, info os.FileInfo) bool {
		return strings.HasPrefix(info.Name(), ".") || info.Name() == "vendor"
	}
	matches, _ := FindFiles(".", acceptFn, skipDirFn, true)
	if len(matches) != 2 {
		t.Fatal("cannot find filewalk.go and filewalk_test.go")
	}
}
