package glob

import (
	"os"
	"testing"
	
	"github.com/subchen/go-stack/assert"
)

func TestIsGlobPattern(t *testing.T) {
	assert.True(t, IsGlobPattern("a*b"))
	assert.True(t, IsGlobPattern("ab/*.txt"))
	assert.True(t, IsGlobPattern("ab/c.tx?"))
	assert.True(t, IsGlobPattern("ab/[a-z].txt"))
	assert.False(t, IsGlobPattern("abc.txt"))
	assert.False(t, IsGlobPattern("ab/c.txt"))
}

func TestAntStyleGlob(t *testing.T) {
	root, err := ioutil.TempDir("", "testdata")
	defer os.RemoveAll(root)
	
	files := []string{
		"t0.txt",
		"a/t1.txt",
		"a/t2.jpg",
		"a/b/t3.txt",
		"x/y/t4.jpg",
	}
	
	for _, file := range files {
		file = filepath.Join(root, file)
		os.MkdirAll(filepath.Dir(file), 0755)
		f, _ := os.CreateFile(file)
		f.Close()
	}

	files := AntStyleGlob(root, []string{"**/*.txt"}, nil, []string{"t1.txt"}, false)
	expected := []string{
		"t0.txt",
		"a/b/t3.txt",
	}
	assert.Equal(t, expected, files)
}
