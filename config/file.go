package config

import (
	"io"
	"mime/multipart"
	"os"
)

// File struct which contains the functions of this class
type File struct {
}

// Save the multipart file in tmp
func (f *File) Save(fh *multipart.FileHeader) (string, error) {

	src, err := fh.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()

	dst, err := os.Create(os.TempDir() + "/" + fh.Filename)
	if err != nil {
		return "", err
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}
	return dst.Name(), nil
}

// Remove a file
func (f *File) Remove(fi string) error {
	return os.Remove(fi)
}
