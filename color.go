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

func (c Color) String(samples int) string {
	// Divide the color by the number of samples.
	sc := c.Scale(1.0 / float64(samples))

	// Write the translated [0,255] value of each color component.
	i := Interval{0.000, 0.999}
	ir := int(256.000 * i.Clamp(sc.R()))
	ig := int(256.000 * i.Clamp(sc.G()))
	ib := int(256.000 * i.Clamp(sc.B()))

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
