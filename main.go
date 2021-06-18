package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

const (
	imageWidth  = 256
	imageHeight = 256
)

func main() {
	var outputFlag string

	flag.StringVar(&outputFlag, "o", "render.png", "the file to output to")
	flag.StringVar(&outputFlag, "output", "render.png", "the file to output to")
	flag.Parse()

	if outputFlag == "" {
		flag.Usage()
		return
	}

	f, err := os.Create(outputFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	topLeft := image.Point{0, 0}
	bottomRight := image.Point{imageWidth, imageHeight}

	img := image.NewRGBA(image.Rectangle{topLeft, bottomRight})
	bounds := img.Bounds()

	// Looping over Y first and X second is more likely to result in better memory access patterns than X first and Y
	// second.
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var (
				r = float64(x) / float64(bounds.Max.X-1)
				g = 1 - float64(y)/float64(bounds.Max.Y-1)
				b = 0.25
			)
			img.Set(x, y, color.RGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), math.MaxUint8})
		}
	}

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}
