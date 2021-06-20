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
	var (
		tmp1 = gort.NewVec3(0, 0, -1)
		tmp2 = gort.NewVec3(1, 1, 1)
	)

	t := hitSphere(tmp1, 0.5, r)
	if t > 0 {
		tmp3 := new(gort.Vec3)
		r.At(t, tmp3)
		tmp3.Sub(tmp3, tmp1)
		tmp3.Unit()
		tmp3.Add(tmp3, tmp2)
		return tmp3.Mul(tmp3, 0.5)
	}

	tmp1.Set(r.Direction)
	tmp1.Unit()
	t = 0.5 * (tmp1.Y + 1)
	tmp2.Mul(tmp2, t)
	tmp1.Set3(0.5, 0.7, 1)
	tmp1.Mul(tmp1, 1-t)
	return tmp1.Add(tmp1, tmp2)
}

func hitSphere(center *gort.Vec3, radius float64, ray *gort.Ray) float64 {
	var (
		oc = new(gort.Vec3).Sub(ray.Origin, center)
		a  = ray.Direction.Dot(ray.Direction)
		b  = 2 * oc.Dot(ray.Direction)
		c  = oc.Dot(oc) - radius*radius
	)
	discriminant := b*b - 4*a*c
	if discriminant > 0 {
		return -1
	}
	return (-b - math.Sqrt(discriminant)) / (2 * a)
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
