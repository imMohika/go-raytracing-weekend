package main

import (
	"fmt"
	"log"
	"os"
	"raytracing-books/geometry"
)

func main() {
	// Image
	aspectRatio := 16.0 / 9.0
	ImageWidth := 400
	ImageHeight := int(float64(ImageWidth) / aspectRatio)

	// File
	file, err := os.Create("./img.ppm")
	if err != nil {
		log.Fatal("Failed to open file")
	}

	if _, err := file.WriteString(
		fmt.Sprintf("P3\n%d %d\n255\n", ImageWidth, ImageHeight)); err != nil {
		log.Fatal("Failed to write to file")
	}

	// Camera
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(ImageWidth) / float64(ImageHeight))
	cameraCenter := geometry.Vec{0, 0, 0}

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := geometry.Vec{viewportWidth, 0, 0}
	viewportV := geometry.Vec{0, -viewportHeight, 0}

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	pixelDeltaU := viewportU.Scale(1.0 / float64(ImageWidth))
	pixelDeltaV := viewportV.Scale(1.0 / float64(ImageHeight))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := cameraCenter.Minus(geometry.Vec{0, 0, focalLength}, viewportU.Scale(0.5), viewportV.Scale(0.5))
	pixel00Loc := viewportUpperLeft.Plus(pixelDeltaU.Plus(pixelDeltaV).Scale(0.5))

	for j := 0; j < ImageHeight; j++ {
		log.Printf("Scanlines remaining: %d\n", ImageHeight-j)
		for i := 0; i < ImageWidth; i++ {
			pixelCenter := pixel00Loc.Plus(pixelDeltaU.Scale(float64(i)), pixelDeltaV.Scale(float64(j)))
			rayDir := pixelCenter.Minus(cameraCenter)
			ray := Ray{cameraCenter, rayDir}
			pixelColor := ray.Color()

			if _, err := file.WriteString(pixelColor.String()); err != nil {
				log.Fatal("Failed to write to file")
			}

			if _, err := file.WriteString("\n"); err != nil {
				log.Fatal("Failed to write to file")
			}
		}
	}

	log.Println("Done.")
}
