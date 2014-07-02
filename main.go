package main

import (
	"./fibersegmentation"
	"fmt"
	"sync"
)

func main() {

	test := fibersegmentation.Project{
		Src:  "./img/thresh.png",
		Dest: "out.png",
		Thresholds: fibersegmentation.ThresholdPair{
			Low:  1.4,
			High: 120,
		},
	}

	test.Init()

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	// convert image, write to disk
	go func() {
		test.SaveSegmentedImage()

		print("Bild gespeichert.\n")
		waitGroup.Done()
	}()

	// analyze fibers etc.
	go func() {
		print("Analysiere Fasern...\n")
		fibers := test.ConnectedComponents()

		for n, fiber := range fibers {
			print("Fläche(", n, "): ", len(fiber), "\n")
		}

		fmt.Printf("%v", fibersegmentation.SortFibers(fibers)[21:32])

		waitGroup.Done()
	}()

	// Let mainthread wait for all threads
	waitGroup.Wait()

}
