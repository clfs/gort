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

	"github.com/clfs/gort"
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
	origin        = &gort.Vec3{X: 0, Y: 0, Z: 0}
	horizontal    = &gort.Vec3{X: viewportWidth, Y: 0, Z: 0}
	vertical      = &gort.Vec3{X: 0, Y: viewportHeight, Z: 0}
	topLeftCorner = &gort.Vec3{
		X: -viewportWidth / 2,
		Y: -viewportHeight / 2,
		Z: focalLength,
	}
)

func rayColor(r *gort.Ray) *gort.Vec3 {
	if hitSphere(&gort.Vec3{X: 0, Y: 0, Z: 1}, 0.5, r) {
		return &gort.Vec3{X: 1, Y: 0, Z: 0}
	}
	unit := r.Direction.Unit()
	t := 0.5 * (unit.Y + 1)
	tmp1 := &gort.Vec3{X: 1, Y: 1, Z: 1}
	tmp2 := &gort.Vec3{X: 0.5, Y: 0.7, Z: 1}
	return tmp1.Add(tmp1.Mul(tmp1, t), tmp2.Mul(tmp2, 1-t))
}

func hitSphere(center *gort.Vec3, radius float64, ray *gort.Ray) bool {
	var (
		oc = new(gort.Vec3).Sub(ray.Origin, center)
		a  = ray.Direction.Dot(ray.Direction)
		b  = 2 * oc.Dot(ray.Direction)
		c  = oc.Dot(oc) - radius*radius
	)
	return b*b-4*a*c > 0
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
			var (
				u = float64(x) / float64(bounds.Max.X-1)
				v = float64(y) / float64(bounds.Max.Y-1)

				r          gort.Ray
				tmp1, tmp2 gort.Vec3
				c          = &color.RGBA{A: math.MaxUint8}
			)

			tmp2.Mul(horizontal, u)
			tmp1.Add(topLeftCorner, &tmp2)
			tmp2.Mul(vertical, v)
			tmp1.Add(&tmp1, &tmp2)
			tmp1.Sub(&tmp1, origin)

			r.Origin = origin
			r.Direction = &tmp1

			img.Set(x, y, rayColor(&r).Color(c))
		}
	}

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}

	bar.Finish()
}
