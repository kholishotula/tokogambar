package imagestrg

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/riandyrn/tokogambar/internal/core/entity"
)

var testCases = []string{
	"image_1.jpg",
	"image_2.jpg",
	"image_3.jpg",
	"image_4.jpg",
	"image_5.jpg",
	"image_6.jpg",
	"image_7.jpg",
	"image_8.jpg",
	"image_9.jpg",
	"image_10.jpg",
	"image_11.jpg",
	"image_12.jpg",
	"image_13.jpg",
	"image_14.jpg",
	"image_15.jpg",
}

func TestGetImages(t *testing.T) {
	basePath := "../../../../static/images"
	infos, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Errorf("unable to read dir due: %w", err)
	}
	var images []entity.Image
	for _, info := range infos {
		if info.IsDir() {
			continue
		}
		images = append(images, entity.Image{
			Filename: info.Name(),
			Hash:     "",
		})
	}

	imageStorage := New(
		Config{
			Images: images,
		},
	)

	for _, image := range imageStorage.GetImages(context.Background()) {
		if stringInSlice(image.Filename, testCases) == false {
			t.Log("Image with file name", image.Filename, "is outside of source")
			t.Fail()
		}
	}
}

func stringInSlice(str string, slice []string) bool {
	for _, element := range slice {
		if element == str {
			return true
		}
	}
	return false
}
