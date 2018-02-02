// Package tar implements the Archive interface providing tar.gz archiving and compression.
package tar

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"os"
)

// Archive as tar.gz
type Archive struct {
	f  *os.File
	gw *gzip.Writer
	tw *tar.Writer
}

// Close all closeables
func (a Archive) Close() error {
	if err := a.tw.Close(); err != nil {
		return err
	}
	if err := a.gw.Close(); err != nil {
		return err
	}
	if err := a.f.Close(); err != nil {
		return err
	}
	return nil
}

// New tar.gz archive
func New(filename string) Archive {
	f, err := os.Create(filename)
	if err != nil {
		panic("unable to create .tar.gz file: " + filename)
	}
	gw := gzip.NewWriter(f)
	tw := tar.NewWriter(gw)
	return Archive{
		f:  f,
		gw: gw,
		tw: tw,
	}
}

// Add file to the archive
func (a Archive) Add(name, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return errors.New("unable to add dir into archive: " + path)
	}

	header := new(tar.Header)
	header.Name = name
	header.Size = stat.Size()
	header.Mode = int64(stat.Mode())
	header.ModTime = stat.ModTime()
	if err := a.tw.WriteHeader(header); err != nil {
		return err
	}
	if _, err := io.Copy(a.tw, file); err != nil {
		return err
	}
	return err
}
