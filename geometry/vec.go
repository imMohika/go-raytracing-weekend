package geometry

import "math"

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

func (v Vec) Plus(other Vec) Vec {
	return Vec{
		v.X() + other.X(),
		v.Y() + other.Y(),
		v.Z() + other.Z(),
	}
}

func (v Vec) Minus(other Vec) Vec {
	return Vec{
		v.X() - other.X(),
		v.Y() - other.Y(),
		v.Z() - other.Z(),
	}
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
