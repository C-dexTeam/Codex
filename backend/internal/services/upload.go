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
}

func newUploadService(
	utilService IUtilService,
) *uploadService {
	return &uploadService{
		utilService: utilService,
		uploadDir:   "uploads",
	}
}

func (s *uploadService) SaveImage(file *multipart.FileHeader, filePath string) error {
	if err := s.checkDirectory(); err != nil {
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

func (s *uploadService) MainDir() string {
	return s.uploadDir
}

func (s *uploadService) DeleteFile(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}
func (s *uploadService) checkDirectory() error {
	if err := file.CheckDir(s.uploadDir); err != nil {
		if err := file.CreateDir(s.uploadDir); err != nil {
			return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusInternalServerError, serviceErrors.ErrCreateDirectoryError)
		}
	}

	return nil
}
