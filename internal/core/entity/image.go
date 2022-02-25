package entity

import (
	"image"

	"github.com/corona10/goimagehash"
)

type Image struct {
	Filename string
	Hash     string
}

func GetImageHash(image image.Image) (string, error) {
	imageHash, err := goimagehash.PerceptionHash(image)
	if err != nil {
		return "", err
	}

	return imageHash.ToString(), nil
}
