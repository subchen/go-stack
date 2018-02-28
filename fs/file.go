package fs

import (
	"io/ioutil"
	"os"
	"time"
)

func FileGetSize(filename string) int64 {
	info, err := os.Stat(filename)
	if err != nil {
		return -1
	}
	return info.Size()
}

func FileGetLastModified(filename string) time.Time {
	info, err := os.Stat(filename)
	if err != nil {
		return time.Unix(0, 0)
	}
	return info.ModTime()
}

func FileGetBytes(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func FileGetString(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func FileWriteBytes(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}

func FileWriteString(filename string, data string) (string, error) {
	return ioutil.WriteFile(filename, []byte(data), 0644)
}
