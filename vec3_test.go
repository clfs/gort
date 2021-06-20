package gort_test

import (
	"math"
	"testing"

	. "github.com/clfs/gort"
)

func approxEqVec3(a, b *Vec3, delta float64) bool {
	return math.Abs(a.X-b.X) < delta &&
		math.Abs(a.Y-b.Y) < delta &&
		math.Abs(a.Z-b.Z) < delta
}

func TestVec3_Cross(t *testing.T) {
	t.Parallel()
	var (
		a    = &Vec3{3, -3, 1}
		b    = &Vec3{4, 9, 2}
		want = &Vec3{-15, -2, 39}
	)
	got := new(Vec3).Cross(a, b)
	if !approxEqVec3(want, got, 1e-8) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestVec3_Cross_Aliased(t *testing.T) {
	t.Parallel()
	var (
		a    = &Vec3{3, -3, 1}
		b    = &Vec3{4, 9, 2}
		want = &Vec3{-15, -2, 39}
	)
	a.Cross(a, b)
	if !approxEqVec3(want, a, 1e-8) {
		t.Errorf("want %v, got %v", want, a)
	}
}

func BenchmarkVec3_Cross(b *testing.B) {
	b.ReportAllocs()
	var (
		r = new(Vec3)
		s = &Vec3{4, 5, 6}
		t = &Vec3{7, 8, 9}
	)
	for i := 0; i < b.N; i++ {
		r.Cross(s, t)
	}
}
