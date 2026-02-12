package S3

import (
	"context"
	"io"
	"log"
	"study-stack/internal/shared/env"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Storage struct {
	client *s3.Client
	bucket string
}

func NewStorage(bucket string) Storage {
	cfg := aws.Config{
		Region:      "us-east-1",
		Credentials: credentials.NewStaticCredentialsProvider(env.Config.S3_KEY, env.Config.S3_SECRET, ""),
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(env.Config.S3_URL)
		o.UsePathStyle = true
	})
	ctx := context.Background()
	if err := ensureBucketExists(ctx, client, bucket); err != nil {
		log.Fatalf("Error creating s3 bucket: %v\n", err)
	}
	return &s3Storage{
		client: client,
		bucket: bucket,
	}
}

func (s *s3Storage) Upload(ctx context.Context, key string, r io.Reader, contentType string) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        r,
		ContentType: aws.String(contentType),
	})
	return err
}

func (s *s3Storage) GetURL(ctx context.Context, key string) (string, error) {
	presignClient := s3.NewPresignClient(s.client)
	req, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func (s *s3Storage) Delete(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	return err
}
