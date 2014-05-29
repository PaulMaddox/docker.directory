package storage

import (
	"io"
	"log"
	"os"

	"github.com/rlmcpherson/s3gof3r"
)

// AwsProvider is a storage provider that allows Docker images
// to be stored on Amazon S3
type AwsProvider struct {
	AccessKey string
	SecretKey string
	Bucket    string
}

// NewAwsProvider creates a new instance of the AWS S3 storage provider
// with the authentication credentials provided, ready for upload/download operations.
func NewAwsProvider(key, secret, bucket string) StorageProvider {
	return &AwsProvider{
		AccessKey: key,
		SecretKey: secret,
		Bucket:    bucket,
	}
}

// Get downloads a file from AWS S3 and provides an io.ReadCloser to read it from
func (s *AwsProvider) Get(path string) (io.ReadCloser, error) {

	s3gof3r.SetLogger(os.Stdout, "", log.LstdFlags, true)
	s3 := s3gof3r.New("", s3gof3r.Keys{
		AccessKey: s.AccessKey,
		SecretKey: s.SecretKey,
	})

	bucket := s3.Bucket(s.Bucket)

	log.Printf("Downloading %s", path)

	r, _, err := bucket.GetReader(path, nil)
	return r, err

}

// Put uploads a to a file in AWS S3 from the provided io.Reader until io.EOF
func (s *AwsProvider) Put(path string, r io.ReadCloser) error {

	s3gof3r.SetLogger(os.Stdout, "", log.LstdFlags, true)
	s3gof3r.DefaultConfig.Md5Check = true

	s3 := s3gof3r.New("", s3gof3r.Keys{
		AccessKey: s.AccessKey,
		SecretKey: s.SecretKey,
	})

	bucket := s3.Bucket(s.Bucket)

	log.Printf("Uploading %s", path)
	w, err := bucket.PutWriter(path, nil, nil)
	if err != nil {
		log.Printf("Error opening S3 bucket %s (%s)", s.Bucket, err)
		return err
	}

	bytes, err := io.Copy(w, r)
	log.Printf("IO copy complete")
	if err != nil {
		log.Printf("Error uploading %s (%s)", path, err)
		return err
	}

	log.Printf("Uploaded %s (%d bytes)", path, bytes)

	if err := r.Close(); err != nil {
		log.Printf("Error closing reader: %s", err)
	}

	if err := w.Close(); err != nil {
		log.Printf("Error closing writer: %s", err)
	}

	return nil

}
