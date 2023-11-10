package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"

)


func main() {

	webcam, err := gocv.VideoCaptureDevice(int(2))

	if err != nil {
		fmt.Printf("Error opening video file: %v\n", err)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Video")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	// Get the FPS property
	fps := webcam.Get(gocv.VideoCaptureFPS)
	fmt.Printf("FPS of the video: %.2f\n", fps)

	// Calculate delay for the desired frame rate
	// delay := int(1000 / fps)
	green_color := color.RGBA{0, 255, 0, 0}

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Println("End of video")
			break
		}
		// Add text to the image
		text := fmt.Sprintf("FPS: %.2f", fps)
		gocv.PutText(&img, text, image.Pt(10, 30), gocv.FontHersheyPlain, 1.2, green_color, 2)

		window.IMShow(img)
		window.WaitKey(1)
		// Introduce a delay to control the frame rate
		// gocv.WaitKey(delay)
	}
}
