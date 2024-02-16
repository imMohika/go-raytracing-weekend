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

	var world HittableList
	world.Add(Sphere{geometry.Vec{0, 0, -1}, 0.5})
	world.Add(Sphere{geometry.Vec{0, -100.5, -1}, 100})

	frame := NewFrame(400, 16/9)
	cam := NewCamera(frame)
	cam.Render(file, world)

	log.Println("Done.")
}
