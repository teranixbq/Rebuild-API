package storage

import (
	"errors"
	"mime/multipart"
	"recything/app/config"
	"recything/utils/constanta"

	supabase "github.com/supabase-community/storage-go"
)

type storageConfig struct {
	sb *supabase.Client
}

type StorageInterface interface {
	Upload(image *multipart.FileHeader) (string, error)
}

func NewStorage(sb *supabase.Client) StorageInterface {
	return &storageConfig{
		sb: sb,
	}
}

func InitStorage(cfg *config.AppConfig) *supabase.Client {
	storageClient := supabase.NewClient(constanta.URL_STORAGE, cfg.API_STORAGE, nil)
	return storageClient
}

var (
	contentType = "image/png"
	bucket      = "recything"
	upsert      = true
)

func (sc *storageConfig) Upload(image *multipart.FileHeader) (string, error) {
	file, err := image.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	result, err := sc.sb.UploadFile(bucket, image.Filename, file, supabase.FileOptions{
		ContentType: &contentType,
		Upsert:      &upsert,
	})
	if err != nil {
		return "", errors.New(err.Error())
	}
	url := constanta.URL + result.Key
	return url, nil
}
