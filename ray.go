package main

import (
	"raytracing-books/geometry"
	"syscall"
)

type Ray struct {
	origin geometry.Vec
	dir    geometry.Vec
}

func (r Ray) Origin() geometry.Vec {
	return r.origin
}

func (r Ray) Dir() geometry.Vec {
	return r.dir
}

func (r Ray) At(t float64) geometry.Vec {
	return r.Origin().Plus(r.dir.Scale(t))
}

func (r Ray) Color(world HittableList) Color {
	if rec := world.Hit(r, Interval{0, syscall.INFINITE}); rec != nil {
		return Color{
			rec.normal.X() + 1,
			rec.normal.Y() + 1,
			rec.normal.Z() + 1,
		}.Scale(0.5)
	}

	a := 0.5 * (r.Dir().Y() + 1.0)
	white := Color{1, 1, 1}.Scale(1.0 - a)
	blue := Color{0.5, 0.7, 1}.Scale(a)
	return white.Plus(blue)
}
