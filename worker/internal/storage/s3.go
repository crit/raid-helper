package storage

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Storage struct {
	session *session.Session
	bucket  *string
}

func NewS3(bucketName, region string) (FileStorage, error) {
	var awsSession, err = session.NewSession(aws.NewConfig().WithRegion(region))
	if err != nil {
		return nil, err
	}

	return &S3Storage{bucket: &bucketName, session: awsSession}, nil
}

func (s *S3Storage) Get(fileName string) ([]byte, error) {
	var data []byte
	writer := aws.NewWriteAtBuffer(data)
	downloader := s3manager.NewDownloader(s.session)

	_, err := downloader.Download(writer, &s3.GetObjectInput{
		Bucket: s.bucket,
		Key:    aws.String(fileName),
	})

	if err != nil {
		return nil, err
	}

	return writer.Bytes(), nil
}

func (s *S3Storage) Set(fileName string, data []byte) error {
	uploader := s3manager.NewUploader(s.session)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      s.bucket,
		Key:         aws.String(fileName),
		ContentType: aws.String(ContentType(fileName)),
		Body:        bytes.NewReader(data),
	})

	return err
}

func (s *S3Storage) SetModel(fileName string, version int, data []byte) error {
	return s.Set(fullName(fileName, version, "json"), data)
}

func (s *S3Storage) GetModel(fileName string, version int) ([]byte, error) {
	return s.Get(fullName(fileName, version, "json"))
}
