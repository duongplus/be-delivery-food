package uploadprovider

import (
	"be-food-delivery/common"
	"context"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
	DeleteFileUploaded(ctx context.Context, dst string) error
}
