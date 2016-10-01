package imagetocsv

import (
	"image"
	"image/color"
	"os"
	"strconv"
	"strings"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func ReadImage(filename string) image.Image {
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
func ColorToBrightness(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return (int)((r + g + b) / 3)
}

func ConvertToCSV(filename string) string {
	img := ReadImage(filename)
	bounds := img.Bounds()
	var vals []string

	if (strings.Contains(filename, "square")) {
		vals = append(vals, "square")
	} else {
		vals = append(vals, "notsquare")
	}

	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			vals = append(vals, strconv.Itoa(ColorToBrightness(img.At(i, j))))
		}
	}

	return strings.Join(vals, ",")
}

func CreateHeaderRow(pixelcount int) string {
	var vals []string;
	vals = append(vals, "label")
	for i := 0; i < pixelcount; i++ {
		vals = append(vals, "pixel" + strconv.Itoa(i))
	}

	return strings.Join(vals, ",")
}