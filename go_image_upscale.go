package main

import (
	"image"
	"image/jpeg"
	"net/http"
	"os"
)

func openImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	i, err := jpeg.Decode(f)
	return i, err
}

type httpImage struct {
	img image.Image
}

func (i httpImage) serveImage(w http.ResponseWriter, r *http.Request) {
	err := jpeg.Encode(w, i.img, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func resize(img image.Image, scale int) image.Image {
	x_max, y_max := img.Bounds().Max.X*scale, img.Bounds().Max.Y*scale
	new_image := image.NewNRGBA(image.Rect(0, 0, x_max, y_max))
	for x := 0; x < x_max; x++ {
		for y := 0; y < y_max; y++ {
			c := img.At(x/scale, y/scale)
			new_image.Set(x, y, c)
		}
	}
	return new_image
}

func main() {
	img, err := openImage("./Go_gopher.jpg")
	if err != nil {
		panic(err)
	}

	new_img := resize(img, 4)

	http_image := httpImage{new_img}
	http.HandleFunc("/", http_image.serveImage)
	http.ListenAndServe(":8080", nil)
}
