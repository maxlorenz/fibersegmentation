package fibersegmentation

import (
	"image"
	"image/png"
	"os"
)

type Project struct {
	Src, Dest  string
	Thresholds ThresholdPair
	Image      image.Image
}

type ThresholdPair struct {
	Low  float32
	High uint8
}

func (self *Project) Init() {
	
	srcF, err := os.Open(self.Src)

	if err != nil {
		panic("Bild nicht gefunden")
	}

	self.Image, _, err = image.Decode(srcF)

	if err != nil {
		panic("Konnte Bild nicht umwandeln")
	}
}

func (self *Project) SaveSegmentedImage() {

	destF, _ := os.OpenFile(self.Dest, os.O_CREATE|os.O_WRONLY, 0666)
	result := Segment(self.Image, self.Thresholds.Low, self.Thresholds.High)

	if err := png.Encode(destF, result); err != nil {
		panic("Konnte Bild nicht speichern")
	}

}
