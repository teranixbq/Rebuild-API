package storage

import (
	"errors"
	"mime/multipart"
	"recything/app/config"
	supabase "github.com/supabase-community/storage-go"
)

var (
	contentType   = "image/png"
	bucket        = "RecyThingAPI"
	upsert        = true
	URL           = "https://vjnyddjlwngtvvndowgo.supabase.co"
	storageClient = supabase.NewClient(URL, config.InitConfig().API_STORAGE, nil)
)

func UploadProof(image *multipart.FileHeader) (string, error) {
	file, err := image.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	result, err := storageClient.UploadFile(bucket, image.Filename, file, supabase.FileOptions{
		ContentType: &contentType,
		Upsert:      &upsert,
	})
	if err != nil {
		return "", errors.New(err.Error())
	}
	url := URL + result.Key
	return url, nil
}


func UploadThumbnail(image *multipart.FileHeader) (string, error) {
	file, err := image.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	result, err := storageClient.UploadFile(bucket, image.Filename, file, supabase.FileOptions{
		ContentType: &contentType,
		Upsert:      &upsert,
	})
	if err != nil {
		return "", errors.New(err.Error())
	}
	url := URL + result.Key
	return url, nil
}
