package main

import (
	"fmt"
	"image"
	"image/draw"
	"os"
	"./terrain"
)

func main() {
	if(len(os.Args) < 2) {
		fmt.Printf("Usage:\n\t%s FILENAME\n", os.Args[0])
		os.Exit(1)
	}

	img, err := terrain.Load(os.Args[1])

	if(err != nil) {
		os.Exit(1)
	}

	out := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y))

	terrain.Equalize(img.(draw.Image))
	terrain.Save(img, "output-eq.png")

    terrain.MapColors(out, img)
	terrain.Save(out, "output.png")
}