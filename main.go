package main

import (
	"log"
	"os"
	"raytracing-books/geometry"
)

func main() {
	file, err := os.Create("./img.ppm")
	if err != nil {
		log.Fatal("Failed to open file")
	}

	// Materials
	matGround := Lambertian{Color{0.8, 0.8, 0}}
	matCenter := Lambertian{Color{0.7, 0.3, 0.3}}
	matLeft := Metal{Color{0.8, 0.8, 0.8}, 0.3}
	matRight := Metal{Color{0.8, 0.6, 0.2}, 1.0}

	var world HittableList
	world.Add(Sphere{geometry.Vec{0, -100.5, -1}, 100, matGround})
	world.Add(Sphere{geometry.Vec{0, 0, -1}, 0.5, matCenter})
	world.Add(Sphere{geometry.Vec{-1, 0, -1}, 0.5, matLeft})
	world.Add(Sphere{geometry.Vec{1, 0, -1}, 0.5, matRight})

	frame := NewFrame(400, 16/9)
	cam := NewCamera(frame, 100, 10)
	cam.Render(file, world)

	log.Println("Done.")
}
