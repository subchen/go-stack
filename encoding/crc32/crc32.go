package crc32

import (
	"hash/crc32"
	"io"
	"os"
)

// SumBytes returns crc32sum of data
func SumBytes(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

// SumString returns crc32sum of data
func SumString(data string) uint32 {
	return crc32.ChecksumIEEE([]byte(data))
}

// SumFile returns crc32sum of file
func SumFile(file string) (uint32, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	h := crc32.NewIEEE()
	if _, err := io.Copy(h, f); err != nil {
		return 0, err
	}

	return h.Sum32(), nil
}
