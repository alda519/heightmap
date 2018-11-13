package terrain

import (
	_ "fmt"
	"image"
	"image/png"
	"image/color"
	"image/draw"
	"log"
	"os"
)

func Load(file string) (image.Image, error) {
	reader, err := os.Open(file)

	if(err != nil) {
		log.Println("cannot load image", err)
		return nil, err
	}

	img, _, err := image.Decode(reader)

	if(err != nil) {
		log.Println("cannot decode image", err)
		return nil, err
	}
	
	return img, nil
}

func Save(img image.Image, file string) (error) {
	writer, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0664)

	if(err != nil) {
		log.Fatal("Cannot write image", err)
	}

	return png.Encode(writer, img)
}


/**
 * histogram equalization
 */
func Equalize(img draw.Image) {
	bounds := img.Bounds()
	min, max := uint32(65535), uint32(0)

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			r, _, _, _ := img.At(x, y).RGBA()

			if r > max {
				max = r
			}
			if r < min {
				min = r
			}
		}
	}

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			r, _, _, _ := img.At(x, y).RGBA()
			c := uint8(float64(r - min) / float64(max - min) * 255)
			img.Set(x, y, color.RGBA{c, c, c, 255})
		}
	}
}

/**
 * Map grayscale image to green..yellow image.
 */
func MapColors(dst draw.Image, img image.Image) {
	bounds := img.Bounds()

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			
			r, _, _, _ := img.At(x, y).RGBA()

			// TODO: need better color mapping function
			ratio := float64(r) / 65536


			//c := color.RGBA{60 + uint8(ratio * 140), 203, 54, 255}
			c := color.RGBA{120 + uint8(ratio * 130), 203, 54, 255}

			dst.Set(x, y, c)
		}
	}
}
