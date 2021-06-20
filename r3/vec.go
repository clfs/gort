package r3

import (
	"image/color"
	"math"
)

// Vec is a 3D vector.
type Vec struct {
	X, Y, Z float64
}

// NewVec returns a new Vec with the given coordinates.
func NewVec(x, y, z float64) Vec {
	return Vec{x, y, z}
}

// Add returns a+b.
func Add(a, b Vec) Vec {
	return Vec{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

// Sub returns a-b.
func Sub(a, b Vec) Vec {
	return Vec{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

// Scale returns v scaled by f.
func Scale(v Vec, f float64) Vec {
	return Vec{
		X: v.X * f,
		Y: v.Y * f,
		Z: v.Z * f,
	}
}

// Mag returns the magnitude of v.
func Mag(v Vec) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Mag2 returns the magnitude squared of v.
func Mag2(v Vec) float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Unit returns the unit vector in the direction of v.
// It panics if v is the zero vector.
func Unit(v Vec) Vec {
	return Scale(v, 1/Mag(v))
}

// Dot returns a dot b.
func Dot(a, b Vec) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Cross returns a cross b.
func Cross(a, b Vec) Vec {
	return Vec{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

// Color returns the color corresponding to v.
// [0, 1] for XYZ is mapped to [0, 255] for RGB.
// The alpha channel is always 255.
func (v Vec) Color() color.RGBA {
	if v.X < 0 || v.Y < 0 || v.Z < 0 || v.X > 1 || v.Y > 1 || v.Z > 1 {
		panic("uhoh")
	}
	return color.RGBA{
		R: uint8(v.X * math.MaxUint8),
		G: uint8(v.Y * math.MaxUint8),
		B: uint8(v.Z * math.MaxUint8),
		A: math.MaxUint8,
	}
}
