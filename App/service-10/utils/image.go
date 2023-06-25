package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"github.com/nfnt/resize"
	"strings"
)

func ResizeImage(base64Image string, width uint,height uint) string {
	imageData, err := decodeBase64Image(base64Image)
	if err != nil{
		return ""
	}
	newImage := resize.Resize(width, height, imageData, resize.Lanczos3)
	return encodeImageToBase64(newImage)
}

func decodeBase64Image(base64Image string) (image.Image, error) {
	// Remove the data URI scheme prefix if it exists
	base64Data := strings.TrimPrefix(base64Image, "data:image/png;base64,")

	// Decode the base64-encoded image data
	// Decode the base64-encoded image data using a buffered reader
	imageReader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Data))

	// Create an image from the decoded bytes
	img, err := png.Decode(imageReader)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func encodeImageToBase64(img image.Image) string {
	buffer := new(bytes.Buffer)
	err := png.Encode(buffer, img)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return ""
	}

	return base64.StdEncoding.EncodeToString(buffer.Bytes())
}