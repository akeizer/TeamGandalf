package imagetocsv

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"
	"strings"
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

	if strings.Contains(filename, "square") {
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
	var vals []string
	vals = append(vals, "label")
	for i := 0; i < pixelcount; i++ {
		vals = append(vals, "pixel"+strconv.Itoa(i))
	}

	return strings.Join(vals, ",")
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	var train = flag.Bool("train", false, "Use the input files as training")
	var help = flag.Bool("h", false, "Display the help information")

	flag.Parse()

	if *help {
		os.Exit(1)
	}
	args := flag.Args()
	outfilename := args[0]

	// open output file
	if !*train {
		outfile, err := os.Create(outfilename)
		if err != nil {
			panic(fmt.Sprintf("Failed to create output file: %s ", err))
		}
		for _, arg := range args[1:] {
			outfile.WriteString(ConvertToCSV(arg))
		}
		defer outfile.Close()
	}
}
