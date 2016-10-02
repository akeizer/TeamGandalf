package imagetocsv

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"
	"strings"
)

func ConvertImageSet(outfilename string, inputFiles ...string) error {
	outfile, err := os.Create(outfilename)
    if err != nil {
        return err
    }
    defer outfile.Close()

    totalpixels := 400

    outfile.WriteString(createHeaderRow(totalpixels) + "\n")
    for _, input := range inputFiles[:] {
    	outString, err := convertToCSV(input)
    	if err != nil {
    		return err
    	}
        outfile.WriteString(outString + "\n")
    }
    return nil
}

func ReadImage(filename string) (image.Image, error) {
	fileReader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(fileReader)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// ignore alpha for now
func colorToBrightness(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return (int)((r + g + b) / 3)
}

func convertToCSV(filename string) (string, error) {
	img, err := ReadImage(filename)
	if err != nil {
		return "", err
	}
	bounds := img.Bounds()
	var vals []string

	if strings.Contains(filename, "square") {
		vals = append(vals, "square")
	} else {
		vals = append(vals, "notsquare")
	}

	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			vals = append(vals, strconv.Itoa(colorToBrightness(img.At(i, j))))
		}
	}

	return strings.Join(vals, ","), nil
}

func createHeaderRow(pixelcount int) string {
	var vals []string
	vals = append(vals, "label")
	for i := 0; i < pixelcount; i++ {
		vals = append(vals, "pixel"+strconv.Itoa(i))
	}

	return strings.Join(vals, ",")
}

