package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"raytracing-books/geometry"
)

type Camera struct {
	frame       Frame
	samples     int // Count of random samples for each pixel
	depth       int // Maximum number of ray bounces into scene
	center      geometry.Vec
	pixelDeltaU geometry.Vec
	pixelDeltaV geometry.Vec
	pixel00Loc  geometry.Vec
}

func NewCamera(frame Frame, samples, depth int, vfov float64, lookFrom, lookAt, vup geometry.Vec) Camera {
	center := lookFrom

	// Determine viewport dimensions.
	focalLength := lookFrom.Minus(lookAt).Length()
	theta := vfov * math.Pi / 180
	h := math.Tan(theta / 2.0)
	viewportHeight := 2.0 * h * focalLength
	viewportWidth := viewportHeight * (float64(frame.width) / float64(frame.height))

	// Calculate the u,v,w unit basis vectors for the camera coordinate frame.
	w := lookFrom.Minus(lookAt).Unit()
	u := vup.Cross(w).Unit()
	v := w.Cross(u)

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := u.Scale(viewportWidth)
	viewportV := v.Inverse().Scale(viewportHeight)

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	pixelDeltaU := viewportU.Scale(1.0 / float64(frame.width))
	pixelDeltaV := viewportV.Scale(1.0 / float64(frame.height))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := center.Minus(w.Scale(focalLength), viewportU.Scale(0.5), viewportV.Scale(0.5))
	pixel00Loc := viewportUpperLeft.Plus(pixelDeltaU.Plus(pixelDeltaV).Scale(0.5))

	return Camera{
		frame,
		samples,
		depth,
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
			pixelColor := Color{0, 0, 0}
			for s := 0; s < c.samples; s++ {
				r := c.getRay(i, j)
				clr := r.Color(world, c.depth)
				//log.Println(pixelColor, clr)
				pixelColor = pixelColor.Plus(clr)
			}

			if _, err := f.WriteString(pixelColor.String(c.samples)); err != nil {
				log.Fatal("Failed to write to file")
			}

			if _, err := f.WriteString("\n"); err != nil {
				log.Fatal("Failed to write to file")
			}
		}
	}
}

func (c Camera) getRay(i int, j int) Ray {
	pixelCenter := c.pixel00Loc.Plus(c.pixelDeltaU.Scale(float64(i)), c.pixelDeltaV.Scale(float64(j)))
	pixelSample := pixelCenter.Plus(c.sampleSquare())

	rayDir := pixelSample.Minus(c.center)
	return Ray{c.center, rayDir}
}

func (c Camera) sampleSquare() geometry.Vec {
	px := -0.5 + rand.Float64()
	py := -0.5 + rand.Float64()
	return c.pixelDeltaU.Scale(px).Plus(c.pixelDeltaV.Scale(py))
}
