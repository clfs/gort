package gort

type Ray struct {
	Origin, Direction *Vec3
}

// At sets dst to the linear interpolation of r at t, and returns dst.
func (r *Ray) At(t float64, dst *Vec3) *Vec3 {
	return dst.Add(dst.Mul(r.Direction, t), r.Origin)
}
