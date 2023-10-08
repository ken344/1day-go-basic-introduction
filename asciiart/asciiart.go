package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"io/ioutil"
	"os"
)

func main() {
	reader, ferr := os.Open("wow.png")
	if ferr != nil {
		fmt.Println("ファイルが開けませんでした", ferr)
		os.Exit(1)
	}

	img, name, derr := image.Decode(reader)
	if derr != nil {
		fmt.Println("画像の変換エラーです", derr)
		os.Exit(1)
	} else {
		fmt.Println("ファイル名:", name)
		fmt.Println("画像のサイズ:", img.Bounds().Size())
	}

	defer reader.Close()

	marks := []string{"*", "+", "-"}
	var marksStr string

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			grayness := c.Y / (255 / 3)
			if c.Y == 255 {
				grayness = 2
			}
			marksStr += marks[grayness]
		}
		marksStr += "\n"
	}
	wdata := []byte(marksStr)
	werr := ioutil.WriteFile("ascii_art.txt", wdata, 0777)
	if werr != nil {
		fmt.Println("ファイルの書き込みエラーです", werr)
		os.Exit(1)
	} else {
		fmt.Println("ファイルを保存しました")
	}

}
