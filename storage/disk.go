package storage

import (
	"io"
	"os"
)

// DiskProvider is a storage provider that stores your Docker images on local disk.
type DiskProvider struct {
	Directory string
}

// NewDiskProvider creates a new disk storage provider instance
func NewDiskProvider() StorageProvider {
	return &DiskProvider{
		Directory: "/tmp/images",
	}
}

// Reader returns an io.ReadCloser for the specified path
func (d *DiskProvider) Reader(path string) (io.ReadCloser, error) {
	return os.Open(d.Directory + string(os.PathSeparator) + path)
}

// Writer returns an io.WriteCloser for the specified path
func (d *DiskProvider) Writer(path string) (io.WriteCloser, error) {
	return os.OpenFile(d.Directory+string(os.PathSeparator)+path, os.O_CREATE|os.O_WRONLY, 0666)
}
