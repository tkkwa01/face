package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image/color"
)

func main() {
	// Webカメラを開く
	webcam, err := gocv.VideoCaptureDevice(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	// 画像ウィンドウを作成
	window := gocv.NewWindow("Face Detection")
	defer window.Close()

	// カスケード分類器を読み込む
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("haarcascade_frontalface_default.xml") {
		fmt.Println("Error loading cascade file")
		return
	}

	// 画像を格納するMatを作成
	img := gocv.NewMat()
	defer img.Close()

	for {
		// フレームをキャプチャ
		if ok := webcam.Read(&img); !ok {
			fmt.Println("Cannot read webcam")
			return
		}
		if img.Empty() {
			continue
		}

		// 顔を検出
		rects := classifier.DetectMultiScale(img)
		for _, r := range rects {
			gocv.Rectangle(&img, r, color.RGBA{0, 0, 255, 0}, 3)
		}

		// ウィンドウに表示
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
