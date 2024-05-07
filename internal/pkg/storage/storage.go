package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type Storage interface {
	Upload(ctx context.Context, filename string, file multipart.File) (string, error)
}

type storage struct {
	minio *minio.Client
}

func New(minioClient *minio.Client) *storage {
	return &storage{minio: minioClient}
}

func (s *storage) Upload(ctx context.Context, filename string, file multipart.File) (string, error) {
	filename = fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(filename))
	uploadInfo, err := s.minio.PutObject(ctx, "tracks", filename, file, 0, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return uploadInfo.Key, nil
}
