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

		waitGroup.Done()
	}()

	// analyze fibers etc.
	go func() {
		fibers := test.ConnectedComponents()

		sorted := fibersegmentation.SortFibers(fibers)

		fmt.Println("Nummer,Länge,Pixel,Krümmung")
		for i, fiber := range sorted {
			l := fibersegmentation.FiberLength(fiber)
			a := len(fiber)
			r := 0.0

			if l == 0 {
				r = 0.0
			} else {
				r = float64(a) / l
			}
			fmt.Printf("%v,%f,%v,%v\n", i, l, a, r)
		}

		waitGroup.Done()
	}()

	// Let mainthread wait for all threads
	waitGroup.Wait()

}
