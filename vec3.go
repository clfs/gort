package gort

import (
	"math"
	"math/rand"
)

type Vec3 struct {
	X, Y, Z float64
}

// Set sets v to a and returns v.
func (v *Vec3) Set(a *Vec3) *Vec3 {
	v.X = a.X
	v.Y = a.Y
	v.Z = a.Z
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

// Unit sets v to the unit vector in the direction of v and returns v.
func (v *Vec3) Unit() *Vec3 {
	return v.Div(v, v.Mag())
}
