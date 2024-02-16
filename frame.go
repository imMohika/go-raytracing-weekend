package main

type Frame struct {
	width, height int
}

func NewFrame(width int, aspectRatio float64) Frame {
	height := int(float64(width) / aspectRatio)

	return Frame{
		width, height,
	}
}
