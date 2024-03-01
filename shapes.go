package main

import (
	"math"
	"raytracing-books/geometry"
)

type Interval struct {
	min, max float64
}

func (i Interval) Contains(x float64) bool {
	return i.min <= x && x <= i.max
}

func (i Interval) Surrounds(x float64) bool {
	return i.min < x && x < i.max
}

func (i Interval) Clamp(x float64) float64 {
	if x < i.min {
		return i.min
	}
	if x > i.max {
		return i.max
	}
	return x
}

type HitRecord struct {
	p         geometry.Vec
	normal    geometry.Vec
	mat       Material
	t         float64
	frontFace bool
}

// MakeHitRecord NOTE: the parameter `outwardNormal` is assumed to have unit length.
func MakeHitRecord(r Ray, p geometry.Vec, mat Material, t float64, outwardNormal geometry.Vec) HitRecord {
	n := outwardNormal.Inverse()
	if r.Dir().Dot(outwardNormal) < 0 {
		// Front face
		n = outwardNormal
	}
	return HitRecord{
		p,
		n,
		mat,
		t,
		true,
	}
}

type Hittable interface {
	Hit(r Ray, t Interval) *HitRecord
}

type Sphere struct {
	center geometry.Vec
	radius float64
	mat    Material
}

func (s Sphere) Hit(r Ray, t Interval) *HitRecord {
	oc := r.Origin().Minus(s.center)

	a := r.Dir().LengthSquared()
	bHalf := oc.Dot(r.Dir())
	c := oc.LengthSquared() - s.radius*s.radius
	discriminant := bHalf*bHalf - a*c
	if discriminant < 0 {
		return nil
	}

	squirted := math.Sqrt(discriminant)

	root := (-bHalf - squirted) / a
	if !t.Surrounds(root) {
		root := (-bHalf + squirted) / a
		if !t.Surrounds(root) {
			return nil
		}
	}

	p := r.At(root)
	outwardNormal := p.Minus(s.center).Scale(1.0 / s.radius).Unit()
	rec := MakeHitRecord(r, p, s.mat, root, outwardNormal)
	return &rec
}
