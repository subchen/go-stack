// Package zip implements the Archive interface providing zip archiving
// and compression.
package zip

import (
	"archive/zip"
	"errors"
	"io"
	"os"
)

// Archive zip struct
type Archive struct {
	f *os.File
	z *zip.Writer
}

// Close all closeables
func (a Archive) Close() error {
	if err := a.z.Close(); err != nil {
		return err
	}
	if err := a.f.Close(); err != nil {
		return err
	}
	return nil
}

// New zip archive
func New(filename string) Archive {
	f, err := os.Create(filename)
	if err != nil {
		panic("unable to create zip file: " + filename)
	}
	return Archive{
		f: f,
		z: zip.NewWriter(f),
	}
}

// Add a file to the zip archive
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

	header, err := zip.FileInfoHeader(stat)
	if err != nil {
		return err
	}
	header.Method = zip.Deflate
	header.Name = name
	f, err := a.z.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, file)
	return err
}
