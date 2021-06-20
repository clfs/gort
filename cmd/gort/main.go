package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/cheggaaa/pb/v3"

	"github.com/clfs/gort/r3"
)

const (
	// Image.
	aspectRatio = 16. / 9
	imageWidth  = 400
	imageHeight = int(imageWidth / aspectRatio)

	// Camera.
	viewportHeight = 2
	viewportWidth  = aspectRatio * viewportHeight
	focalLength    = 1
)

var (
	// Camera.
	origin        = r3.NewVec(0, 0, 0)
	horizontal    = r3.NewVec(viewportWidth, 0, 0)
	vertical      = r3.NewVec(0, viewportHeight, 0)
	topLeftCorner = r3.NewVec(-viewportWidth/2, -viewportHeight/2, focalLength)
)

func rayColor(r r3.Ray) color.RGBA {
	t := hitSphere(r3.NewVec(0, 0, 1), 0.5, r)
	if t > 0 {
		n := r3.Sub(r3.At(r, t), r3.NewVec(0, 0, 1))
		return r3.Scale(
			r3.Add(n, r3.NewVec(1, 1, 1)),
			0.5,
		).Color()
	}

	// If the sphere wasn't hit.
	t = 0.5 * (r3.Unit(r.Direction).Y + 1)
	return r3.Add(
		r3.Scale(r3.NewVec(1, 1, 1), t),
		r3.Scale(r3.NewVec(.5, .7, 1), 1-t),
	).Color()
}

func hitSphere(center r3.Vec, radius float64, r r3.Ray) float64 {
	var (
		oc    = r3.Sub(r.Origin, center)
		a     = r3.Mag2(r.Direction)
		halfB = r3.Dot(oc, r.Direction)
		c     = r3.Mag2(oc) - radius*radius
	)
	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return -1
	}
	return (-halfB - math.Sqrt(discriminant)) / a
}

func main() {
	log.SetFlags(0)

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

	bar := pb.StartNew(bounds.Max.Y)

	// Looping over Y first and X second is more likely to result in better memory access patterns than X first and Y
	// second.
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		bar.Increment()
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			u := float64(x) / float64(bounds.Max.X)
			v := float64(y) / float64(bounds.Max.Y)
			r := r3.Ray{
				Origin: origin,
				Direction: r3.Add(
					topLeftCorner,
					r3.Add(
						r3.Scale(horizontal, u),
						r3.Sub(
							r3.Scale(vertical, v),
							origin,
						),
					)),
			}
			img.Set(x, y, rayColor(r))
		}
	}

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}

	bar.Finish()
}
