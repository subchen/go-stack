package md5

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// SumBytes returns md5sum of data
func SumBytes(data []byte) string {
	sum := md5.Sum(data)
	return fmt.Sprintf("%x", sum)
}

// SumString returns md5sum of data
func SumString(data string) string {
	sum := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", sum)
}

// SumFile returns md5sum of file
func SumFile(file string) (string, error) {
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

// GenerateSumFile generates a .md5 file
func GenerateSumFile(file string) error {
	sum, err := SumFile(file)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file+".md5", []byte(sum), 0644)
}
