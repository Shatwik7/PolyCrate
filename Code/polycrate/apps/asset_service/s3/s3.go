package asset

// import (
// 	"bytes"
// 	"context"
// 	"log"
// 	"time"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/config"
// 	s3Presign "github.com/aws/aws-sdk-go-v2/feature/s3/presign"
// 	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
// )

// type S3Service struct {
// 	Client     *s3.Client
// 	Presigner  *s3Presign.PresignClient
// 	BucketName string
// }

// func NewS3Service(bucket string) *S3Service {
// 	cfg, err := config.LoadDefaultConfig(context.TODO())
// 	if err != nil {
// 		log.Fatalf("unable to load AWS SDK config: %v", err)
// 	}
// 	client := s3.NewFromConfig(cfg)
// 	presigner := s3Presign.NewPresignClient(client)

// 	return &S3Service{
// 		Client:     client,
// 		Presigner:  presigner,
// 		BucketName: bucket,
// 	}
// }

// func (s *S3Service) UploadFile(key string, content []byte) error {
// 	_, err := s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
// 		Bucket: aws.String(s.BucketName),
// 		Key:    aws.String(key),
// 		Body:   bytes.NewReader(content),
// 	})
// 	return err
// }

// func (s *S3Service) DeleteFile(key string) error {
// 	_, err := s.Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
// 		Bucket: aws.String(s.BucketName),
// 		Key:    aws.String(key),
// 	})
// 	return err
// }

// func (s *S3Service) GetPresignedURL(key string, expiry time.Duration) (string, error) {
// 	req, err := s.Presigner.PresignGetObject(context.TODO(), &s3.GetObjectInput{
// 		Bucket: aws.String(s.BucketName),
// 		Key:    aws.String(key),
// 	}, s3Presign.WithExpires(expiry))
// 	if err != nil {
// 		return "", err
// 	}
// 	return req.URL, nil
// }
