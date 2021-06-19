package gort

type Point3 struct {
	X, Y, Z float64
}

// Add sets p to the sum a+b and returns p.
func (p *Point3) Add(a, b *Point3) *Point3 {
	p.X = a.X + b.X
	p.Y = a.Y + b.Y
	p.Z = a.Z + b.Z
	return p
}
