package storage

import (
	"bytes"
	"io"
	"log"
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

// Get pretends to download a file, but actually just serves back dummy data
func (d *DummyProvider) Get(path string) (io.ReadCloser, error) {
	return &ClosingBuffer{
		bytes.NewBufferString("Dummy storage provider doesn't care about your data"),
	}, nil
}

// Put pretends to upload a file, but actually just casually accepts and never delivers
func (d *DummyProvider) Put(path string, r io.Reader, length int64) error {
	log.Printf("DummyProvider: Uploading file %s...", path)
	return nil
}
