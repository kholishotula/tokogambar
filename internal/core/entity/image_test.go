package entity

import (
	"image/jpeg"
	"os"
	"testing"
)

var testCases = []struct {
	filename string
	hash     string
}{
	{
		filename: "image_1.jpg",
		hash:     "p:ae463458d2658f76",
	},
	{
		filename: "image_2.jpg",
		hash:     "p:81c53337c3c72576",
	},
	{
		filename: "image_3.jpg",
		hash:     "p:f819b9e62a9d9d20",
	},
	{
		filename: "image_4.jpg",
		hash:     "p:9cbec64191bb2e4a",
	},
	{
		filename: "image_5.jpg",
		hash:     "p:87c63f86703cb1b1",
	},
	{
		filename: "image_6.jpg",
		hash:     "p:d185d14cd6b21cb7",
	},
	{
		filename: "image_7.jpg",
		hash:     "p:9a89b8f0a6747656",
	},
	{
		filename: "image_8.jpg",
		hash:     "p:94933a9667ae9171",
	},
	{
		filename: "image_9.jpg",
		hash:     "p:92b3721ea965c696",
	},
	{
		filename: "image_10.jpg",
		hash:     "p:c9b7701f74cd2521",
	},
}

func TestGetImageHash(t *testing.T) {
	for _, testCase := range testCases {
		file, _ := os.Open("../../../static/images/" + testCase.filename)
		defer file.Close()

		img, _ := jpeg.Decode(file)
		hash, _ := GetImageHash(img)

		if testCase.hash != hash {
			t.Log(
				"The perception hash result of",
				testCase.filename,
				"is",
				hash,
				". It should be",
				testCase.hash,
			)
			t.Fail()
		}
	}
}
