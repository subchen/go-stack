package checksum

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Sha256sum(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func Sha256sumAsFile(file string) error {
	sum, err := Sha256sum(file)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file+".sha256", []byte(sum), 0644)
}
