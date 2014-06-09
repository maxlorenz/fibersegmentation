package main

import (
	"../src/segment"
	"../src/view"
	"image"
	"image/png"
	"log"
	"os"
)

func main() {

	view.Run()

	srcF, err := os.Open("thresh.png")
	destF, err := os.OpenFile("out.png", os.O_CREATE | os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal("Bild nicht gefunden")
	}

	src, _, err := image.Decode(srcF)
	if err != nil {
		log.Fatal("Konnte Bild nicht umwandeln")
	}
	
	result := segment.Segment(src, 1.4, 120)

	if err = png.Encode(destF, result); err != nil {
		log.Fatal("Konnte Bild nicht speichern")
	}

}