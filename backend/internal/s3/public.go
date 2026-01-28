package S3

import "context"

type Storage interface {
	Upload(ctx context.Context, key string, data []byte) error
	GetURL(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}
