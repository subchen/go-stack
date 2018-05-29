package findup

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindFileInCWD(t *testing.T) {
	path, _ := Find("findup.go")
	if path != fullpath("findup.go") {
		t.Fatalf("file not found")
	}
}

func TestFindFileInCWD_NotFound(t *testing.T) {
	path, _ := Find("not-found")
	if path != "" {
		t.Fatalf("file was found")
	}
}

func TestFindGlobFileInCWD(t *testing.T) {
	path, _ := Find("findup.*")
	if path != fullpath("findup.go") {
		t.Fatalf("file not found")
	}
}

func TestFindGlobFileInUpDir(t *testing.T) {
	tmp := ".tmp/a/b/c"
	_ = os.MkdirAll(tmp, 0755)
	defer os.RemoveAll(tmp)

	path, _ := FindInDir(tmp, "findup.go")
	if path != fullpath("findup.go") {
		t.Fatalf("file was found")
	}
}

func fullpath(filename string) string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, filename)
}
