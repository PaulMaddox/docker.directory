package storage

import "io"

const (
	// ProviderAws is a storage provider that interfaces to Amazon S3
	ProviderAws int = 0

	// ProviderDummy is a storage provider that is downright dishonest
	// and doesn't give a damn about your data. It is however extremly fast.
	ProviderDummy int = 1
)

// StorageProvider interface provides various upload/download functions
type StorageProvider interface {
	Get(path string) (io.ReadCloser, error)
	Put(path string, r io.Reader, length int64) error
}
