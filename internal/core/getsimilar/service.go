package getsimilar

import (
	"context"
	"image"

	"github.com/riandyrn/tokogambar/internal/core/entity"
	"github.com/riandyrn/tokogambar/internal/core/levenshtein"
)

type Service interface {
	GetSimilarImages(ctx context.Context, image image.Image) ([]entity.SimilarImage, error)
}

type service struct {
	imageStorage ImageStorage
}

func (s *service) GetSimilarImages(ctx context.Context, image image.Image) ([]entity.SimilarImage, error) {
	hashStr, err := entity.GetImageHash(image)
	if err != nil {
		return nil, err
	}

	simImages := []entity.SimilarImage{}
	for _, image := range s.imageStorage.GetImages(ctx) {
		// change hash similarity checking by using levenshtein distance
		// strings are considered similar if their levenshtein distance is less than or equal to 5
		if levenshtein.DistanceTwoStrings(image.Hash, hashStr) <= 5 {
			simImages = append(simImages, entity.SimilarImage{
				Filename:        image.Filename,
				SimilarityScore: 100.0,
			})
		}
	}
	return simImages, nil
}

type ServiceConfig struct {
	ImageStorage ImageStorage
}

func NewService(cfg ServiceConfig) Service {
	s := &service{
		imageStorage: cfg.ImageStorage,
	}
	return s
}
