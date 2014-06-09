package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
)

const (

	// ProviderAws is a storage provider that interfaces to Amazon S3
	ProviderAws int = 0

	// ProviderDummy is a storage provider that is downright dishonest
	// and doesn't give a damn about your data. It is however extremly fast.
	ProviderDummy int = 1

	// ProviderDisk is a storage provider that interfaces to local disk
	ProviderDisk int = 2
)

// StorageProvider is responsible for obtaining/managing a reader/writer to
// a storage type (eg disk/s3)
type StorageProvider interface {
	Reader(path string) (io.ReadCloser, error)
	Writer(path string) (io.WriteCloser, error)
}

// Get retrieves a file from the provided storage provider
func Get(path string, provider StorageProvider) (io.ReadCloser, error) {
	return provider.Reader(path)
}

// Put writes a file to the provided storage provider and returns a SHA256 checksum
func Put(path string, r io.ReadCloser, provider StorageProvider) (string, error) {

	writer, err := provider.Writer(path)
	if err != nil {
		return "", err
	}

	// Read out the data from the input reader
	// and write it to the file and also to a SHA256 checksummer
	hash := sha256.New()

	buffer := make([]byte, 1024)
	for {

		_, err := r.Read(buffer)
		if err == io.EOF {
			break
		}

		writer.Write(buffer)
		hash.Write(buffer)

	}

	checksum := hex.EncodeToString(hash.Sum(nil))

	if err := writer.Close(); err != nil {
		log.Printf("Error while closing file %s (%s)", path, err)
		return "", err
	}

	if err := r.Close(); err != nil {
		log.Printf("Error while closing reader %s (%s)", err)
		return "", err
	}

	//log.Printf("Stored %s (%d bytes) to disk", path, written)
	return checksum, nil

}
