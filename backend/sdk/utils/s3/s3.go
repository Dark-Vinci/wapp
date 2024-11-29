package s3

import (
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3 struct {
	region string
	sess   *session.Session
}

type MediaStore interface {
	Upload(f *os.File, key, bucket string) (*string, error)
	Get(key, bucket string) (*string, error)
}

func NewS3(region string) *S3 {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(region)}))

	return &S3{region: region, sess: sess}
}

func (s *S3) Get(key, bucket string) (*string, error) {
	svc := s3.New(s.sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		log.Println("Failed to sign request", err)
		return nil, err
	}

	return &urlStr, err
}

func (s *S3) Upload(f *os.File, key, bucket string) (*string, error) {
	uploader := s3manager.NewUploader(s.sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   f,
	})

	if err != nil {
		return nil, err
	}

	return &result.UploadID, nil
}
