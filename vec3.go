package gort

import (
	"image/color"
	"math"
	"math/rand"
)

type Vec3 struct {
	X, Y, Z float64
}

// NewVec3 returns a new *Vec3.
func NewVec3(x, y, z float64) *Vec3 {
	return &Vec3{X: x, Y: y, Z: z}
}

// Set sets v's values to a's values and returns v.
func (v *Vec3) Set(a *Vec3) *Vec3 {
	v.X = a.X
	v.Y = a.Y
	v.Z = a.Z
	return v
}

// Set3 sets v's values and returns v.
func (v *Vec3) Set3(x, y, z float64) *Vec3 {
	v.X = x
	v.Y = y
	v.Z = z
	return v
}

// Rand sets v to a random vector and returns v.
func (v *Vec3) Rand() *Vec3 {
	v.X = rand.Float64()
	v.Y = rand.Float64()
	v.Z = rand.Float64()
	return v
}

// Add sets v to the sum a+b and returns v.
func (v *Vec3) Add(a, b *Vec3) *Vec3 {
	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	v.Z = a.Z + b.Z
	return v
}

// Sub sets v to the difference a-b and returns v.
func (v *Vec3) Sub(a, b *Vec3) *Vec3 {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	v.Z = a.Z - b.Z
	return v
}

// Mul sets v to the scalar multiplication a*b and returns v.
func (v *Vec3) Mul(a *Vec3, b float64) *Vec3 {
	v.X = a.X * b
	v.Y = a.Y * b
	v.Z = a.Z * b
	return v
}

// Div sets v to the scalar division a/b and returns v.
func (v *Vec3) Div(a *Vec3, b float64) *Vec3 {
	v.X = a.X / b
	v.Y = a.Y / b
	v.Z = a.Z / b
	return v
}

// Scale scales v by f and returns v.
func (v *Vec3) Scale(f float64) *Vec3 {
	v.X *= f
	v.Y *= f
	v.Z *= f
	return v
}

// Mag returns the magnitude of v.
func (v *Vec3) Mag() float64 {
	return math.Sqrt(v.MagSquared())
}

// MagSquared returns the magnitude squared of v.
func (v *Vec3) MagSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Dot returns the dot product of v and a.
func (v *Vec3) Dot(a *Vec3) float64 {
	return v.X*a.X + v.Y*a.Y + v.Z*a.Z
}

// Cross sets v to the cross product of a and b and returns v.
func (v *Vec3) Cross(a *Vec3, b *Vec3) *Vec3 {
	vX := a.Y*b.Z - a.Z*b.Y
	vY := a.Z*b.X - a.X*b.Z
	vZ := a.X*b.Y - a.Y*b.X
	v.X = vX
	v.Y = vY
	v.Z = vZ
	return v
}

func Cross2(a, b Vec3) Vec3 {
	return Vec3{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

// Unit sets v to the unit vector in the direction of v and returns v.
func (v *Vec3) Unit() *Vec3 {
	return v.Div(v, v.Mag())
}

// Color maps the XYZ coordinates of v onto the RGB channels of dst,
// then returns dst. The alpha channel isn't updated.
func (v *Vec3) Color(dst *color.RGBA) *color.RGBA {
	// [0.0, 1.0] is mapped to [0, 255].
	// Both underflow and overflow can potentially occur.
	dst.R = uint8(v.X * math.MaxUint8)
	dst.G = uint8(v.Y * math.MaxUint8)
	dst.B = uint8(v.Z * math.MaxUint8)
	return dst
}
