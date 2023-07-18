package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func SaveImage(file *multipart.FileHeader, filename string) error {
	storagePath := viper.GetString("Image.Source")
	targetPath := filepath.Join(storagePath, filename)
	target, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer target.Close()

	uploadFile, err := file.Open()
	if err != nil {
		return err
	}
	defer uploadFile.Close()

	_, err = io.Copy(target, uploadFile)
	if err != nil {
		return nil
	}
	return nil
}
