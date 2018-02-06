package fs

import (
	"io/ioutil"
	"os"
	"time"
)

func fileSize(file string) int64 {
	info, err := os.Stat(file)
	if err != nil {
		return -1
	}
	return info.Size()
}

func fileLastModified(file string) time.Time {
	info, err := os.Stat(file)
	if err != nil {
		return time.Unix(0, 0)
	}
	return info.ModTime()
}

func fileGetBytes(file string) []byte {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
	return data
}

func fileGetString(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}
	return string(data)
}
