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

func (r Ray) Color(world HittableList, depth int) Color {
	if b, rec := world.Hit(r, Interval{0.001, syscall.INFINITE}); b {
		// If we've exceeded the ray bounce limit, no more light is gathered.
		if depth <= 0 {
			return Color{0, 0, 0}
		}

		if b, ray, attenuation := rec.mat.Scatter(r, rec); b {
			return ray.Color(world, depth-1).Multiply(attenuation)
		}
		return Color{0, 0, 0}
	}

	a := 0.5 * (r.Dir().Y() + 1.0)
	white := Color{1, 1, 1}.Scale(1.0 - a)
	blue := Color{0.5, 0.7, 1}.Scale(a)
	return white.Plus(blue)
}
