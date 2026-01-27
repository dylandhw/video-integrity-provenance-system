package main

import (
	"gocv.io/x/gocv"
)

var image = "./sample.jpg"

// temporary image processing
func main() {
	window := gocv.NewWindow("sample image for testing")
	windowB := gocv.NewWindow("sample image for testing 2.0")

	sampleImg := gocv.IMRead(image, gocv.IMReadColor)
	sampleGray := gocv.NewMat()

	gocv.CvtColor(sampleImg, &sampleGray, gocv.ColorBGRToGray)

	window.IMShow(sampleImg)
	windowB.IMShow(sampleGray)
	windowB.WaitKey(0)

}
