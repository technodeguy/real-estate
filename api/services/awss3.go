package services

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AwsS3Service struct {
	accessKeyId     string
	secretAccessKey string
	bucketName      string
	s3              *s3.S3
}

func NewAwsS3Service(accessKeyId, secretAccessKey, bucketName string) *AwsS3Service {
	return &AwsS3Service{accessKeyId, secretAccessKey, bucketName, nil}
}

func (a *AwsS3Service) Initialize() {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(a.accessKeyId, a.secretAccessKey, ""),
	})

	if err != nil {
		log.Fatalf("Error by connecting to aws", err)
	}

	a.s3 = s3.New(sess)
}

func (a *AwsS3Service) GetPresignedUrl(mimeType, filename string) (map[string]string, error) {

	req, _ := a.s3.PutObjectRequest(&s3.PutObjectInput{
		Bucket:      aws.String(a.bucketName),
		Key:         aws.String(filename),
		ContentType: aws.String(mimeType),
	})

	str, err := req.Presign(15 * time.Minute)

	if err != nil {
		return map[string]string{}, err
	}

	return map[string]string{
		"key":      filename,
		"mimeType": mimeType,
		"url":      str,
	}, nil
}
