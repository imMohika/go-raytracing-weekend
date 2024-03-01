package geometry

import (
	"math"
	"math/rand"
)

type Vec [3]float64

func (v Vec) X() float64 {
	return v[0]
}

func (v Vec) Y() float64 {
	return v[1]
}

func (v Vec) Z() float64 {
	return v[2]
}

func (v Vec) Inverse() Vec {
	return Vec{
		-v.X(),
		-v.Y(),
		-v.Z(),
	}
}

func (v Vec) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec) LengthSquared() float64 {
	return v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z()
}

func (v Vec) Plus(vecs ...Vec) Vec {
	x := v.X()
	y := v.Y()
	z := v.Z()
	for _, vec := range vecs {
		x += vec.X()
		y += vec.Y()
		z += vec.Z()
	}
	return Vec{x, y, z}
}

func (v Vec) Minus(vecs ...Vec) Vec {
	x := v.X()
	y := v.Y()
	z := v.Z()
	for _, vec := range vecs {
		x -= vec.X()
		y -= vec.Y()
		z -= vec.Z()
	}
	return Vec{x, y, z}
}

func (v Vec) Multiply(other Vec) Vec {
	return Vec{
		v.X() * other.X(),
		v.Y() * other.Y(),
		v.Z() * other.Z(),
	}
}

func (v Vec) Scale(n float64) Vec {
	return Vec{
		v.X() * n,
		v.Y() * n,
		v.Z() * n,
	}
}

func (v Vec) Divide(other Vec) Vec {
	return Vec{
		v.X() / other.X(),
		v.Y() / other.Y(),
		v.Z() / other.Z(),
	}
}

func (v Vec) Dot(other Vec) float64 {
	return v.X()*other.X() + v.Y()*other.Y() + v.Z()*other.Z()
}

func (v Vec) Cross(other Vec) Vec {
	return Vec{
		v.Y()*other.Z() - v.Z()*other.Y(),
		v.Z()*other.X() - v.X()*other.Z(),
		v.X()*other.Y() - v.Y()*other.X(),
	}
}

func (v Vec) Unit() Vec {
	return v.Scale(1.0 / v.Length())
}

func RandVec() Vec {
	return Vec{rand.Float64(), rand.Float64(), rand.Float64()}
}

func RandBoundedVec(min, max float64) Vec {
	return Vec{
		min + (max-min)*rand.Float64(),
		min + (max-min)*rand.Float64(),
		min + (max-min)*rand.Float64(),
	}
}

func RandVecInUnitSphere() Vec {
	for {
		p := RandBoundedVec(-1, 1)
		if p.LengthSquared() < 1.0 {
			return p
		}
	}
}

func RandUnitVec() Vec {
	return RandVecInUnitSphere().Unit()
}

func RandVecOnHemisphere(normal Vec) Vec {
	onSphere := RandUnitVec()
	if onSphere.Dot(normal) > 0.0 {
		// In the same hemisphere as the normal
		return onSphere
	}

	return onSphere.Inverse()
}

func (v Vec) NearZero() bool {
	// Return true if the vector is close to zero in all dimensions.
	z := 1e-8
	return (math.Abs(v.X()) < z) && (math.Abs(v.Y()) < z) && (math.Abs(v.Z()) < z)
}

func Reflect(v, n Vec) Vec {
	return v.Minus(n.Scale(2.0 * v.Dot(n)))
}

func Refract(uv, n Vec, etai float64) Vec {
	cos := math.Min(uv.Inverse().Dot(n), 1.0)
	outPrep := n.Scale(cos).Plus(uv).Scale(etai)
	outParallel := n.Scale(-math.Sqrt(math.Abs(1.0 - outPrep.LengthSquared())))
	return outPrep.Plus(outParallel)
}
