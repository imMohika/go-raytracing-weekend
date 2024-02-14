package main

import (
	"fmt"
	"raytracing-books/geometry"
)

type Color geometry.Vec

func (c Color) R() float64 {
	return c[0]
}

func (c Color) G() float64 {
	return c[1]
}

func (c Color) B() float64 {
	return c[2]
}

func (c Color) String() string {
	ir := int(255.999 * c.R())
	ig := int(255.999 * c.G())
	ib := int(255.999 * c.B())
	return fmt.Sprintf("%d %d %d", ir, ig, ib)
}
