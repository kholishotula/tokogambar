package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/corona10/goimagehash"
	"github.com/riandyrn/tokogambar/levenshtein"
)

const addr = ":8080"

func main() {
	// initialize database
	dbRecords, err := loadDB()
	if err != nil {
		log.Fatalf("unable to initialize database due: %v", err)
	}
	// attach handler
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./images"))))
	http.HandleFunc("/similars", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			WriteAPIResp(w, NewErrorResp(NewErrNotFound()))
			return
		}
		// parse request body
		var rb searchReqBody
		err := json.NewDecoder(r.Body).Decode(&rb)
		if err != nil {
			WriteAPIResp(w, NewErrorResp(NewErrBadRequest(err.Error())))
			return
		}
		// validate request body
		err = rb.Validate()
		if err != nil {
			WriteAPIResp(w, NewErrorResp(err))
			return
		}
		// search similar images
		imgData, err := rb.GetByte()
		if err != nil {
			WriteAPIResp(w, NewErrorResp(err))
			return
		}

		// change similar image searching method by using perceptual hash

		// old
		// similarImages, err := searchSimilarImages(dbRecords, imgData)
		// if err != nil {
		// 	WriteAPIResp(w, NewErrorResp(err))
		// 	return
		// }

		// new
		img, _, err := image.Decode(bytes.NewReader(imgData))
		if err != nil {
			WriteAPIResp(w, NewErrorResp(err))
			return
		}
		similarImages, err := searchSimilarImages(dbRecords, img)
		if err != nil {
			WriteAPIResp(w, NewErrorResp(err))
			return
		}

		// output success response
		WriteAPIResp(w, NewSuccessResp(similarImages))
	})
	// start server
	log.Printf("server is listening on %v", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("unable to start server due: %v", err)
	}
}

func loadDB() ([]dbRecord, error) {
	basePath := "./images"
	infos, err := ioutil.ReadDir(basePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read dir due: %w", err)
	}
	var filenames []string
	for _, info := range infos {
		if info.IsDir() {
			continue
		}
		filenames = append(filenames, info.Name())
	}
	var dbRecords []dbRecord
	for _, filename := range filenames {
		// old
		// b, err := ioutil.ReadFile(basePath + "/" + filename)

		// new
		file, err := os.Open(basePath + "/" + filename)
		if err != nil {
			continue
		}
		defer file.Close()

		imageFile, err := jpeg.Decode(file)
		if err != nil {
			continue
		}

		imageHash, err := getImageHash(imageFile)
		if err != nil {
			continue
		}
		dbRecords = append(dbRecords, dbRecord{
			FileName: filename,
			Hash:     imageHash,
		})
	}
	return dbRecords, nil
}

func searchSimilarImages(dbRecords []dbRecord, img image.Image) ([]similarImage, error) {
	// old
	// hashStr := getHash(data)

	// new
	hashStr, err := getImageHash(img)
	if err != nil {
		return nil, err
	}
	simImages := []similarImage{}
	for _, record := range dbRecords {
		// change hash similarity checking by using levenshtein distance
		// strings are considered similar if their levenshtein distance is less than or equal to 5

		// old
		// if record.Hash == hashStr {
		// 	simImages = append(simImages, similarImage{
		// 		FileName:        record.FileName,
		// 		SimilarityScore: 100.0,
		// 	})
		// }

		//new
		if levenshtein.DistanceTwoStrings(record.Hash, hashStr) <= 5 {
			simImages = append(simImages, similarImage{
				FileName:        record.FileName,
				SimilarityScore: 100.0,
			})
		}
	}
	return simImages, nil
}

// function to get image hash (in string) using phash

// old
// func getHash(data []byte) string {
// 	h := sha256.New()
// 	h.Write(data)

// 	return hex.EncodeToString(h.Sum(nil))
// }

// new
func getImageHash(image image.Image) (string, error) {
	imageHash, err := goimagehash.PerceptionHash(image)
	if err != nil {
		return "", err
	}

	return imageHash.ToString(), nil
}
