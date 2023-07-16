package main

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

func main() {
	dirs, _ := filepath.Glob("./video/frames/*.png")
	fmt.Println(len(dirs))
	for _, element := range dirs {
		//fmt.Println(element)
		read_img(element)
	}

}

func read_img(path string) (int, int, int) {
	img_file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	defer img_file.Close()
	img_data, err := png.Decode(img_file)
	if err != nil {
		// Handle error
		print(err.Error())
	}

	img_data = resize.Resize(960, 720, img_data, resize.Lanczos3)
	bounds := img_data.Bounds()
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Loop through all Pixels in the Image
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, _, _, _ := img_data.At(x, y).RGBA()

			str := fmt.Sprint(0, " ")
			if r/256 == 255 {
				str = fmt.Sprint(1, " ")
			}

			file.WriteString(str)
		}
		file.WriteString("\n")
	}

	err = file.Sync()

	return 0, 0, 0
}
