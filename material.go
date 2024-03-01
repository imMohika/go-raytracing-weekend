package main

import (
	"math"
	"math/rand"
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

type Dielectic struct {
	ir float64 // Index of Refraction
}

func (d Dielectic) Scatter(ray Ray, rec HitRecord) (bool, Ray, Color) {
	refractionRatio := d.ir
	if rec.frontFace {
		refractionRatio = 1.0 / d.ir
	}

	unitDir := ray.Dir().Unit()
	cos := math.Min(unitDir.Inverse().Dot(rec.normal), 1.0)
	sin := math.Sqrt(1.0 - cos*cos)
	canRefract := refractionRatio*sin > 1.0
	var dir geometry.Vec
	if !canRefract || d.reflectance(cos, refractionRatio) > rand.Float64() {
		dir = geometry.Reflect(unitDir, rec.normal)
	} else {
		dir = geometry.Refract(unitDir, rec.normal, refractionRatio)
	}

	return true, Ray{rec.p, dir}, Color{1.0, 1.0, 1.0}
}

func (d Dielectic) reflectance(cos, ref float64) float64 {
	// Use Schlick's approximation for reflectance.
	r0 := (1 - ref) / (1 + ref)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}
