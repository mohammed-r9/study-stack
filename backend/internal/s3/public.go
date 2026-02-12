package S3

import (
	"context"
	"io"
)

type Storage interface {
	Upload(ctx context.Context, key string, r io.Reader, contentType string) error
	GetURL(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}
