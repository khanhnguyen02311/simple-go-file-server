package utils

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type Storage struct {
	Type             string
	UploadAuth       bool
	DownloadAuth     bool
	AuthEndpoint     string
	AllowedMIMETypes []string
	MaxFileSize      int
}

func (s *Storage) Init() {
	s.Type = *ConfigArgs.Type
	s.UploadAuth = *ConfigArgs.UploadAuth == "true"
	s.DownloadAuth = *ConfigArgs.DownloadAuth == "true"
	s.AuthEndpoint = *ConfigArgs.AuthEndpoint
	s.AllowedMIMETypes = strings.Split(*ConfigArgs.AllowedMIMETypes, ",")
	s.MaxFileSize = *ConfigArgs.MaxFileSize
	//fmt.Printf("Storage initialized: %+v", s)
}

func (s *Storage) GetFile(fileName string) ([]byte, error) {
	if s.DownloadAuth && s.validatePermission("token", fileName) == false {
		return nil, errors.New("invalid permission")
	}
	switch s.Type {
	case "local":
		return s.getFromLocal(fileName)
	case "s3":
		return s.getFromS3(fileName)
	default:
		return nil, errors.New("invalid storage type")
	}
}

func (s *Storage) UploadFile(file *multipart.FileHeader) (string, error) {
	if s.UploadAuth && s.validatePermission("token", file.Filename) == false {
		return "", errors.New("invalid permission")
	}
	switch s.Type {
	case "local":
		return s.saveToLocal(file)
	case "s3":
		return s.saveToS3(file)
	default:
		return "", errors.New("invalid storage type")
	}
}

func (s *Storage) validatePermission(accessToken string, fileName string) bool {
	fmt.Println("Validating connection...")
	// TODO: ask the application if the token is valid, for later
	return true
}

func (s *Storage) getFromLocal(filename string) ([]byte, error) {
	return nil, nil
}
func (s *Storage) saveToLocal(file *multipart.FileHeader) (string, error) {
	// checking MIME type
	if s.AllowedMIMETypes[0] != "*" {
		if contains(s.AllowedMIMETypes, file.Header.Get("Content-Type")) == false {
			return "", errors.New("invalid MIME type")
		}
	}
	// checking file size
	if s.MaxFileSize != 0 && float64(file.Size)/1024/1024 > float64(s.MaxFileSize) {
		return "", errors.New("file size exceeded. Files must be smaller than " + fmt.Sprint(s.MaxFileSize) + " MB")
	}

	fmt.Println(file.Size)

	source, err := file.Open()
	if err != nil {
		return "", err
	}
	defer source.Close()

	uuidNow, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	fileExt := path.Ext(file.Filename)
	uuidName := uuidNow.String() + fileExt
	destination, err := os.Create("files/public/" + uuidName)
	if err != nil {
		return "", err
	}
	defer destination.Close()
	if _, err = io.Copy(destination, source); err != nil {
		return "", err
	}
	// TODO: create background task to convert the file type to jpeg and move it to the public directory
	return uuidName, nil
}

func (s *Storage) getFromS3(filename string) ([]byte, error) {
	return nil, nil
}

func (s *Storage) saveToS3(file *multipart.FileHeader) (string, error) {
	return "", nil
}
