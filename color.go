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

func (c Color) Scale(n float64) Color {
	return Color{
		c.R() * n,
		c.G() * n,
		c.B() * n,
	}
}

func (c Color) Plus(colors ...Color) Color {
	r := c.R()
	g := c.G()
	b := c.B()
	for _, color := range colors {
		r += color.R()
		g += color.G()
		b += color.B()
	}
	return Color{r, g, b}
}
