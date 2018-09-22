package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)
import _ "image/jpeg"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// How to read images
	var buff bytes.Buffer
	img, err := os.Open("test.jpg")
	failOnError(err, "testing")
	pic, _, err := image.Decode(img)
	failOnError(err, "image decode")
	png.Encode(&buff, pic)
	encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())
	fmt.Println(encodedString)

}
