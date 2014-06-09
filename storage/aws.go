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

// Reader returns an io.ReadCloser for the specified path
func (s *AwsProvider) Reader(path string) (io.ReadCloser, error) {

	s3gof3r.SetLogger(os.Stdout, "", log.LstdFlags, true)
	s3 := s3gof3r.New("", s3gof3r.Keys{
		AccessKey: s.AccessKey,
		SecretKey: s.SecretKey,
	})

	bucket := s3.Bucket(s.Bucket)
	r, _, err := bucket.GetReader(path, nil)
	return r, err

}

// Writer returns an io.WriteCloser for the specified path
func (s *AwsProvider) Writer(path string) (io.WriteCloser, error) {

	s3gof3r.SetLogger(os.Stdout, "", log.LstdFlags, true)
	s3gof3r.DefaultConfig.Md5Check = true

	s3 := s3gof3r.New("", s3gof3r.Keys{
		AccessKey: s.AccessKey,
		SecretKey: s.SecretKey,
	})

	bucket := s3.Bucket(s.Bucket)

	log.Printf("Uploading %s", path)
	return bucket.PutWriter(path, nil, nil)

}
