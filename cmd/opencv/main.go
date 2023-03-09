package main

import "gocv.io/x/gocv"

func main() {

	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		println("Error opening capture device")
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Detector")

	defer window.Close()

}
