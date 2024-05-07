package minIO

import (
	"fmt"
	"net/http"
	"time"

	"github.com/atadzan/audio-library/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func checkStorageHealth(params config.MinIO) error {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(fmt.Sprintf("http://%s/minio/health/live", params.Endpoint))
	if err != nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("can't connect to storage. Error: %v", err)
	}
	return nil
}

func NewMinioClient(cfg config.MinIO) (*minio.Client, error) {
	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyId, cfg.SecretAccessKeyId, ""),
		Secure: false,
		Region: "us-east-1",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create minio client. err: %w", err)
	}

	if err = checkStorageHealth(cfg); err != nil {
		return nil, err
	}

	return minioClient, nil
}
