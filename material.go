package main

import (
	"math"
	"raytracing-books/geometry"
)

type Material interface {
	Scatter(ray Ray, rec HitRecord) (bool, Ray, Color)
}

type Lambertian struct {
	albedo Color
}

func (l Lambertian) Scatter(ray Ray, rec HitRecord) (bool, Ray, Color) {
	dir := rec.normal.Plus(geometry.RandUnitVec())
	if dir.NearZero() {
		dir = rec.normal
	}
	r := Ray{rec.p, dir}
	return true, r, l.albedo
}

type Metal struct {
	albedo Color
	fuzz   float64
}

func (m Metal) Scatter(ray Ray, rec HitRecord) (bool, Ray, Color) {
	reflected := geometry.Reflect(ray.Dir().Unit(), rec.normal)
	r := Ray{rec.p, reflected.Plus(geometry.RandUnitVec().Scale(math.Min(m.fuzz, 1.0)))}
	return true, r, m.albedo
}
