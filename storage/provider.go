package storage

import "io"

const (
	PROVIDER_AWS   int = 0
	PROVIDER_DISK  int = 1
	PROVIDER_DUMMY int = 2
)

// StorageProvider interface provides various upload/download functions
type StorageProvider interface {
	Get(path string) (io.ReadCloser, error)
	Put(path string, r io.Reader, length int64) error
}
