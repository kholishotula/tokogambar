package getsimilar

import (
	"context"
	"image/jpeg"
	"os"
	"testing"

	"github.com/riandyrn/tokogambar/internal/core/entity"
)

var testCases = []struct {
	inputFile   string
	similarFile []string
}{
	{
		inputFile:   "input_1.jpg",
		similarFile: []string{"image_13.jpg"},
	},
	{
		inputFile:   "input_2.jpg",
		similarFile: []string{"image_10.jpg"},
	},
	{
		inputFile:   "input_3.jpg",
		similarFile: []string{"image_13.jpg"},
	},
}

func TestGetSimilarImages(t *testing.T) {
	// initialize new service
	ns := initNewService()

	basePath := "../../../static/input"

	// check for every testcase
	for _, testCase := range testCases {
		inputFile, err := os.Open(basePath + "/" + testCase.inputFile)
		if err != nil {
			continue
		}
		defer inputFile.Close()

		imageFile, err := jpeg.Decode(inputFile)
		if err != nil {
			continue
		}

		// check whether the results of service are equal to the testcase
		simImages, err := ns.GetSimilarImages(context.Background(), imageFile)
		for _, img := range simImages {
			if !stringInSlice(img.Filename, testCase.similarFile) {
				t.Log(
					"Image with file name",
					img.Filename,
					"should not the similar image as",
					testCase.inputFile,
				)
				t.Fail()
			}
		}
	}
}

type mockImageStorage struct {
	images []entity.Image
}

func (is *mockImageStorage) GetImages(ctx context.Context) []entity.Image {
	return is.images
}

func newMockImageStorage() *mockImageStorage {
	return &mockImageStorage{
		images: []entity.Image{
			{
				Filename: "image_10.jpg",
				Hash:     "p:c9b7701f74cd2521",
			},
			{
				Filename: "image_11.jpg",
				Hash:     "p:e4a2c71b39ec0375",
			},
			{
				Filename: "image_12.jpg",
				Hash:     "p:b5350d45e07f0db4",
			},
			{
				Filename: "image_13.jpg",
				Hash:     "p:d5d5ce90426e58da",
			},
			{
				Filename: "image_14.jpg",
				Hash:     "p:9383a8cc4e13b5f3",
			},
			{
				Filename: "image_15.jpg",
				Hash:     "p:d8c1673838c63977",
			},
		},
	}
}

func initNewService() Service {
	cfg := ServiceConfig{
		ImageStorage: newMockImageStorage(),
	}
	ns := NewService(cfg)
	return ns
}

func stringInSlice(str string, slice []string) bool {
	for _, element := range slice {
		if element == str {
			return true
		}
	}
	return false
}
