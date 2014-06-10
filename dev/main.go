package main

import (
	"../src/segment"
	// "../src/view"
	"image"
	"image/png"
	"os"
)

func main() {

	// view.Run()

	srcF, err := os.Open("thresh.png")
	destF, err := os.OpenFile("out.png", os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		panic("Bild nicht gefunden")
	}

	src, _, err := image.Decode(srcF)
	if err != nil {
		panic("Konnte Bild nicht umwandeln")
	}

	result := segment.Segment(src, 1.4, 120)
	analytics := segment.ReadToMemory(src, 1.4, 120)

	print(len(analytics.Fibers), " zu ", src.Bounds().Max.X * src.Bounds().Max.Y)

	if err = png.Encode(destF, result); err != nil {
		panic("Konnte Bild nicht speichern")
	}

}
