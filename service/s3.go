package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const routeTemplate = "https://storage.yandexcloud.net/%s/%s"

type s3ServiceImpl struct {
	client *s3.Client
	bucket string
}

func newS3Service(client *s3.Client, bucket string) *s3ServiceImpl {
	return &s3ServiceImpl{client, bucket}
}

func (s *s3ServiceImpl) UploadFile(uploader uint, filename string, file multipart.File) (string, error) {
	name := fmt.Sprintf("%d-%d-%s", uploader, time.Now().UnixNano(), filename)
	_, err := s.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(name),
		Body:   file,
	})
	return fmt.Sprintf(routeTemplate, s.bucket, name), err
}
