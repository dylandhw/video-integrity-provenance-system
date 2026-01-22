package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Printf("couldn't grab webcam", err)
		return
	}

	defer webcam.Close()

	window := gocv.NewWindow("gocv webcam capture")
	img := gocv.NewMat()
	defer img.Close()

	for {

		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read from webcam")
			break
		}
		window.IMShow(img)

		if window.WaitKey(1) >= 0 {
			break
		}
	}
	fmt.Println("capture finished")
}
