package storage

import (
	"io"

	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

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
	return nil, nil
}

// Put uploads a to a file in AWS S3 from the provided io.Reader until io.EOF
func (s *AwsProvider) Put(path string, r io.Reader, length int64) error {

	auth := aws.Auth{
		AccessKey: s.AccessKey,
		SecretKey: s.SecretKey,
	}

	awss3 := s3.New(auth, aws.USEast)
	bucket := awss3.Bucket(s.Bucket)

	return bucket.PutReader(path, r, length, "text/plain", s3.BucketOwnerFull)

}
