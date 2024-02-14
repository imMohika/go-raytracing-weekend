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

	_, err = file.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", ImageWidth, ImageHeight))
	if err != nil {
		log.Fatal("Failed to write to file")
	}

	for j := 0; j < ImageHeight; j++ {
		log.Printf("Scanlines remaining: %d\n", ImageHeight-j)
		for i := 0; i < ImageWidth; i++ {
			r := float32(i) / (ImageWidth - 1)
			g := float32(j) / (ImageHeight - 1)
			b := float32(0)

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			_, err = file.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
			if err != nil {
				log.Fatal("Failed to write to file")
			}
		}
	}

	log.Println("Done.")
}
