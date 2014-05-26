package storage

import (
	"bytes"
	"io"
	"log"
)

type DummyProvider struct {
}

type ClosingBuffer struct {
	*bytes.Buffer
}

// NewDummyProvider creates a new storage provider instance that
// is downright dishonest and doesn't give a damn about your data
func NewDummyProvider() StorageProvider {
	return &DummyProvider{}
}

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
