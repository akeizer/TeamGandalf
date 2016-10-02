package imagegen

import (
	"fmt"
	"math/rand"
    "os/exec"
)

const script = "gen_images.sh"

func GenerateImages(imagecount int) {
	_, err := exec.Command("/bin/bash", script).Output()
	if err != nil {
        panic(fmt.Sprintf("Failed to create images: %s ", err))
	}
}

func drawString(shape string, size int, x1 int, y1 int, x2 int, y2 int) string {
	switch (shape) {
	case "triangle":
		tx1 := rand.Intn(size/4) + x1 + (size - size/4)
		ty1 := rand.Intn(size) + y1
		tx2 := rand.Intn(size/2) + x1
		ty2 := rand.Intn(size/4) + y1 + (size - size/4)
		tx3 := rand.Intn(size/2) + x1
		ty3 := rand.Intn(size/4) + y1
		return fmt.Sprintf("polygon %d,%d %d,%d, %d,%d", tx1, ty1, tx2, ty2, tx3, ty3)
	case "circle":
		cx := (x1 + x2) / 2
		cy := (y1 + y2) / 2
		return fmt.Sprintf("circle %d,%d %d,%d", cx, cy, cx, y1)
	case "square":
		return fmt.Sprintf("rectangle %d,%d %d,%d", x1, y1, x2, y2)
	}
	panic(fmt.Sprintf("drawString: Unrecognized shape %s", shape))
}

const image_size = 20

func GenerateImage(shape string, filename string) {
	size := 4 + rand.Intn(12)
	x1 := rand.Intn(image_size - size + 1)
	y1 := rand.Intn(image_size - size + 1)
	x2 := x1 + size
	y2 := y1 + size
	var draw = drawString(shape, size, x1, y1, x2, y2)

	_, err := exec.Command("convert",
		"-size", fmt.Sprintf("%dx%d", image_size, image_size),
		"canvas:white",
		"-stroke", "black",
		"-strokewidth", "1",
		"-fill", "black",
		"-draw", draw,
		filename
	).Output()

	if err != nil {
        panic(fmt.Sprintf("Failed to create images: %s ", err))
	}
}
