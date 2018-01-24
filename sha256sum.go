package gstack

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func sha256sumString(file string) (string, error) {
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

func sha256sumFile(file string) error {
	sum, err := sha256sumString(file)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file+".sha256", []byte(sum), 0644)
}
