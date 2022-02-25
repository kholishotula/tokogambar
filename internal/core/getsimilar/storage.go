package getsimilar

import (
	"context"

	"github.com/riandyrn/tokogambar/internal/core/entity"
)

type ImageStorage interface {
	GetImages(ctx context.Context) []entity.Image
}
