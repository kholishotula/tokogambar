package imagestrg

import (
	"context"

	"github.com/riandyrn/tokogambar/internal/core/entity"
)

type Storage struct {
	images []entity.Image
}

func (s *Storage) GetImages(ctx context.Context) []entity.Image {
	return s.images
}

type Config struct {
	Images []entity.Image
}

func New(cfg Config) *Storage {
	return &Storage{images: cfg.Images}
}
