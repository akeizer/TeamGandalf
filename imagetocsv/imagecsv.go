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

func convertToCSV(filename string) string {
	img := ReadImage(filename)
	bounds := img.Bounds()
	var vals []string

	for i := bounds.Min.X; i <= bounds.Max.X; i++ {
		for j := bounds.Min.Y; j <= bounds.Max.Y; j++ {
			vals = append(vals, strconv.Itoa(ColorToBrightness(img.At(i, j))))
		}
	}
	return strings.Join(vals, ",")
}

func main() {
    if (len(os.Args) < 2) {
        helpText()
        os.Exit(1)
    }

    var train = flag.Bool("train", false, "Use the input files as training")
    var help = flag.Bool("h", false, "Display the help information")

    flag.Parse()

    if *help {
        helpText()
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
            outfile.WriteString(imagetocsv.ConvertToCSV(arg))
        }
        defer outfile.Close()
    }
