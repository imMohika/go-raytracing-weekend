package main

import "raytracing-books/geometry"

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

func (r Ray) Color() Color {
	unitDir := r.Dir().Unit()
	a := 0.5 * (unitDir.Y() + 1.0)
	white := Color{1, 1, 1}.Scale(1.0 - a)
	blue := Color{0.5, 0.7, 1}.Scale(a)
	return white.Plus(blue)
}
