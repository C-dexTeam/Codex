package services

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/C-dexTeam/codex/pkg/file"
)

type uploadService struct {
	utilService IUtilService
	uploadDir   string
	courseDir   string
	web3Dir     string
}

func newUploadService(
	utilService IUtilService,
) *uploadService {
	return &uploadService{
		utilService: utilService,
		uploadDir:   "uploads",
		courseDir:   "uploads/courses",
		web3Dir:     "uploads/web3",
	}
}

func (s *uploadService) SaveImage(file *multipart.FileHeader, filePath string) error {
	if err := s.createDirectories(); err != nil {
		return err
	}

	// Dosya uzantısını kontrol ediyoruz
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrInvalidFileType)
	}

	// Dosyayı açıyoruz
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Kaydetmek için hedef dosyayı oluşturuyoruz
	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Dosya içeriğini kopyalıyoruz
	_, err = io.Copy(outFile, src)
	if err != nil {
		return err
	}

	return nil
}

func (s *uploadService) createDirectories() error {
	if err := file.CheckDir(s.uploadDir); err != nil {
		if err := file.CreateDir(s.uploadDir); err != nil {
			return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusInternalServerError, serviceErrors.ErrCreateDirectoryError)
		}
	}
	if err := file.CheckDir(s.courseDir); err != nil {
		if err := file.CreateDir(s.courseDir); err != nil {
			return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusInternalServerError, serviceErrors.ErrCreateDirectoryError)
		}
	}
	if err := file.CheckDir(s.web3Dir); err != nil {
		if err := file.CreateDir(s.web3Dir); err != nil {
			return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusInternalServerError, serviceErrors.ErrCreateDirectoryError)
		}
	}

	return nil
}

func (s *uploadService) UploadDir() string {
	return s.uploadDir
}

func (s *uploadService) CourseDir() string {
	return s.courseDir
}

func (s *uploadService) Web3Dir() string {
	return s.web3Dir
}

func (s *uploadService) DeleteFile(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}
