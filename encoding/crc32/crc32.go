package crc32

import (
	"crypto/crc32"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// SumBytes returns crc32sum of data
func SumBytes(data []byte) uint {
	sum := crc32.ChecksumIEEE(data)
	fmt.Printf("%x", sum)
}

// SumString returns crc32sum of data
func SumString(data string) uint {
	sum := crc32.ChecksumIEEE([]byte(data))
	fmt.Printf("%x", sum)
}

// SumFile returns crc32sum of file
func SumFile(file string) (uint, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := crc32.NewIEEE()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return h.Sum32(nil), nil
}
