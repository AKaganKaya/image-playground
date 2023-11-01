package helpers

import (
	"bytes"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

func ConvertToGrayscale(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	grayImage := image.NewGray(img.Bounds())
	draw.Draw(grayImage, grayImage.Bounds(), img, image.Point{}, draw.Src)

	buffer := new(bytes.Buffer)
	err = jpeg.Encode(buffer, grayImage, nil)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
