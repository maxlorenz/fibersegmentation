package main

import (
	"./fibersegmentation"
	"image"
	"image/png"
	"os"
	"sync"
)

func main() {

	srcF, err := os.Open("./img/thresh.png")
	if err != nil {
		panic("Bild nicht gefunden")
	}

	destF, _ := os.OpenFile("out.png", os.O_CREATE|os.O_WRONLY, 0666)

	src, _, err := image.Decode(srcF)
	if err != nil {
		panic("Konnte Bild nicht umwandeln")
	}

	// Let mainthread wait to finish threads
	wg := sync.WaitGroup{}
	wg.Add(2)

	// convert image, write to disk
	go func() {
		result := fibersegmentation.Segment(src, 1.4, 120)
		if err = png.Encode(destF, result); err != nil {
			panic("Konnte Bild niccht speichern")
		}

		print("Bild gespeichert.\n")

		wg.Done()
	}()

	go func() {
		analytics := fibersegmentation.ReadToMemory(src, 1.4, 120)
		print(len(analytics.Fibers), " zu ", src.Bounds().Max.X*src.Bounds().Max.Y, " Pixel.\n")
		print(len(analytics.Table))

		wg.Done()
	}()

	wg.Wait()

}
