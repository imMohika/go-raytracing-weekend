package main

import (
	"fmt"
	"log"
	"os"
)

const ImageWidth = 256
const ImageHeight = 256

func main() {
	file, err := os.Create("./img.ppm")
	if err != nil {
		log.Fatal("Failed to open file")
	}

	if _, err := file.WriteString(
		fmt.Sprintf("P3\n%d %d\n255\n", ImageWidth, ImageHeight)); err != nil {
		log.Fatal("Failed to write to file")
	}

	for j := 0; j < ImageHeight; j++ {
		log.Printf("Scanlines remaining: %d\n", ImageHeight-j)
		for i := 0; i < ImageWidth; i++ {
			pixel := Color{
				float64(i) / (ImageWidth - 1),
				float64(j) / (ImageHeight - 1),
				float64(0),
			}

			if _, err := file.WriteString(pixel.String()); err != nil {
				log.Fatal("Failed to write to file")
			}

			if _, err := file.WriteString("\n"); err != nil {
				log.Fatal("Failed to write to file")
			}
		}
	}

	log.Println("Done.")
}
