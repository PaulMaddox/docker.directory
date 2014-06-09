package storage

import (
	"bytes"
	"io"
)

// DummyProvider is a storage provider that is downright dishonest
// and doesn't give a damn about your data. It is however, extremly fast.
type DummyProvider struct {
}

// ClosingBuffer is a bytes.Buffer based buffer that implements the Closer interface
type ClosingBuffer struct {
	*bytes.Buffer
}

// NewDummyProvider creates a new dummy storage provider instance
func NewDummyProvider() StorageProvider {
	return &DummyProvider{}
}

// Close closes a buffer
func (cb *ClosingBuffer) Close() error {
	return nil
}

// Reader returns an io.ReadCloser for the specified path
func (d *DummyProvider) Reader(path string) (io.ReadCloser, error) {
	return &ClosingBuffer{
		bytes.NewBufferString("Dummy storage provider doesn't care about your data"),
	}, nil
}

// Writer returns an io.WriteCloser for the specified path
func (d *DummyProvider) Writer(path string) (closer io.WriteCloser, err error) {
	return closer, nil
}
