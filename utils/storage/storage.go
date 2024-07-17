package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	configApp "recything/app/config"
	"recything/utils/constanta"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type storageConfig struct {
	sb *s3.Client
}

type StorageInterface interface {
	Upload(image *multipart.FileHeader) (string, error)
}

func NewStorage(sb *s3.Client) StorageInterface {
	return &storageConfig{
		sb: sb,
	}
}

func InitStorage(cfg *configApp.AppConfig) *s3.Client {

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", cfg.ACCOUNTID),
		}, nil
	})

	config, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.ACCESID, cfg.SECRETKEY, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(config)

	return client
}

var allowedContentTypes = map[string]string{
	".jpg": "image/jpeg",
	".png": "image/png",
	".mp4": "video/mp4",
}

func (sc *storageConfig) Upload(image *multipart.FileHeader) (string, error) {
	godotenv.Load()

	bucket := os.Getenv("BUCKET")

	file, err := image.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	extension := strings.ToLower(filepath.Ext(image.Filename))
	contentType, ok := allowedContentTypes[extension]
	if !ok {
		return "", errors.New("error : format file tidak diizinkan (unauthorized file format)")
	}

	randomKey := uuid.New().String()
	_, err = sc.sb.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &randomKey,
		ContentType: &contentType,
		Body:        file,
	})
	if err != nil {
		return "", err
	}

	publicURL := constanta.URL + randomKey
	return publicURL, nil
}
