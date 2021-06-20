package r3

type Ray struct {
	Origin, Direction Vec
}

// At returns the linear interpolation of r at t.
func At(r Ray, t float64) Vec {
	return Add(r.Origin, Scale(r.Direction, t))
}
