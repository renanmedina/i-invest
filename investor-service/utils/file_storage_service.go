package utils

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type FileStorageService interface {
	Upload(sourceFile string, destFile string) (string, error)
}

type S3FileStorageService struct {
	awsClient *s3manager.Uploader
	configs   *Configs
}

func (s *S3FileStorageService) Upload(sourceFile string, destPath string) (string, error) {
	file, err := os.Open(sourceFile)

	if err != nil {
		return "", err
	}

	defer file.Close()

	uploadOutput, err := s.awsClient.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.configs.AWS_ANNOUNCEMENTS_FILES_BUCKET_NAME),
		Key:    aws.String(destPath),
		Body:   file,
	})

	if err != nil {
		return "", err
	}

	return uploadOutput.Location, nil
}

func NewS3FileStorage() (*S3FileStorageService, error) {
	configs := GetConfigs()
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(configs.AWS_REGION),
	})

	if err != nil {
		return nil, err
	}

	return &S3FileStorageService{
		awsClient: s3manager.NewUploader(session),
		configs:   configs,
	}, nil
}
