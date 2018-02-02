package sha256

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// SumBytes returns sha256sum of data
func SumBytes(data []byte) string {
	sum := sha256.Sum256(data)
	return fmt.Sprintf("%x", sum)
}

// SumString returns sha256sum of data
func SumString(data string) string {
	sum := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", sum)
}

// SumFile returns sha256sum of file
func SumFile(file string) (string, error) {
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

// GenerateSumFile generates a .sha256 file
func GenerateSumFile(file string) error {
	sum, err := SumFile(file)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file+".sha256", []byte(sum), 0644)
}
