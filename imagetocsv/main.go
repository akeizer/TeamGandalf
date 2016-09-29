package imagetocsv

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"strconv"
	"strings"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func readImage(filename string) image.Image {
	fileReader, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(fileReader)
	if err != nil {
		panic(err)
	}
	return img
}

// ignore alpha for now
func colorToBrightness(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return (int)((r + g + b) / 3)
}

func imageToCSV(filename string) string {
	img := readImage(filename)
	bounds := img.Bounds()
	var vals []string

	for i := bounds.Min.X; i <= bounds.Max.X; i++ {
		for j := bounds.Min.Y; j <= bounds.Max.Y; j++ {
			vals = append(vals, strconv.Itoa(colorToBrightness(img.At(i, j))))
		}
	}
	return strings.Join(vals, ",")
}
