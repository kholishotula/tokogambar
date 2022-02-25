package main

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/riandyrn/tokogambar/internal/core/entity"
	"github.com/riandyrn/tokogambar/internal/core/getsimilar"
	"github.com/riandyrn/tokogambar/internal/driven/storage/imagestrg"
	"github.com/riandyrn/tokogambar/internal/driver/rest"

	"github.com/joho/godotenv"
)

func main() {
	// initialize images storage
	basePath := "./static/images"
	infos, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Errorf("unable to read dir due: %w", err)
	}
	var filenames []string
	for _, info := range infos {
		if info.IsDir() {
			continue
		}
		filenames = append(filenames, info.Name())
	}
	var images []entity.Image
	for _, filename := range filenames {
		file, err := os.Open(basePath + "/" + filename)
		if err != nil {
			continue
		}
		defer file.Close()

		imageFile, err := jpeg.Decode(file)
		if err != nil {
			continue
		}

		imageHash, err := entity.GetImageHash(imageFile)
		if err != nil {
			continue
		}
		images = append(images, entity.Image{
			Filename: filename,
			Hash:     imageHash,
		})
	}

	imageStorage := imagestrg.New(
		imagestrg.Config{
			Images: images,
		},
	)
	log.Printf("build image storage successfully")

	// initialize get similar image service
	getsimilarService := getsimilar.NewService(
		getsimilar.ServiceConfig{
			ImageStorage: imageStorage,
		},
	)
	log.Printf("build get-similar-image service successfully")

	// initialize rest api
	api := rest.NewAPI(rest.APIConfig{
		GetSimilarService: getsimilarService,
	})
	log.Printf("build api successfully")

	// load env
	if err = godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	addr := os.Getenv("PORT")

	// initialize server
	server := &http.Server{
		Addr:    ":" + addr,
		Handler: api.Handler(),
	}

	// run server
	log.Printf("server is listening on %v...", addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("unable to start server due: %v", err)
	}
}
