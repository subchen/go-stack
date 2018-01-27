package checksum

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Md5sum(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func Md5sumAsFile(file string) error {
	sum, err := Md5sum(file)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file+".md5", []byte(sum), 0644)
}
