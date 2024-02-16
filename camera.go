package main

import (
	"fmt"
	"io"
	"log"
	"raytracing-books/geometry"
)

type Camera struct {
	frame       Frame
	center      geometry.Vec
	pixelDeltaU geometry.Vec
	pixelDeltaV geometry.Vec
	pixel00Loc  geometry.Vec
}

func NewCamera(frame Frame) Camera {
	center := geometry.Vec{0, 0, 0}

	// Determine viewport dimensions.
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(frame.width) / float64(frame.height))

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := geometry.Vec{viewportWidth, 0, 0}
	viewportV := geometry.Vec{0, -viewportHeight, 0}

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	pixelDeltaU := viewportU.Scale(1.0 / float64(frame.width))
	pixelDeltaV := viewportV.Scale(1.0 / float64(frame.height))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := center.Minus(geometry.Vec{0, 0, focalLength}, viewportU.Scale(0.5), viewportV.Scale(0.5))
	pixel00Loc := viewportUpperLeft.Plus(pixelDeltaU.Plus(pixelDeltaV).Scale(0.5))

	return Camera{
		frame,
		center,
		pixelDeltaU,
		pixelDeltaV,
		pixel00Loc,
	}
}

func (c Camera) Render(f io.StringWriter, world HittableList) {
	if _, err := f.WriteString(
		fmt.Sprintf("P3\n%d %d\n255\n", c.frame.width, c.frame.height)); err != nil {
		log.Fatal("Failed to write to file")
	}

	for j := 0; j < c.frame.height; j++ {
		log.Printf("Scanlines remaining: %d\n", c.frame.height-j)
		for i := 0; i < c.frame.width; i++ {
			pixelCenter := c.pixel00Loc.Plus(c.pixelDeltaU.Scale(float64(i)), c.pixelDeltaV.Scale(float64(j)))
			rayDir := pixelCenter.Minus(c.center)
			ray := Ray{c.center, rayDir}
			pixelColor := ray.Color(world)

			if _, err := f.WriteString(pixelColor.String()); err != nil {
				log.Fatal("Failed to write to file")
			}

			if _, err := f.WriteString("\n"); err != nil {
				log.Fatal("Failed to write to file")
			}
		}
	}
}