package main

import (
	"image"
	"image/png"
	"log"
	"os"
	"./segment"
)

func main() {

	srcF, err := os.Open("images/thresh.png")
	destF, err := os.OpenFile("out.png", os.O_CREATE | os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal("Bild nicht gefunden")
	}

	src, _, err := image.Decode(srcF)
	if err != nil {
		log.Fatal("Konnte Bild nicht umwandeln")
	}
	
	test := new (segment.SegmentImage)
	test.Original = src
	test.Segment(1.4, 120)

	if err = png.Encode(destF, test.Segmented); err != nil {
		log.Fatal(err)
	}

}