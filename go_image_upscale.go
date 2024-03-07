package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strconv"
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

func integerScale(img image.Image, scale int) image.Image {
	x_max, y_max := img.Bounds().Max.X*scale, img.Bounds().Max.Y*scale
	new_img := image.NewNRGBA(image.Rect(0, 0, x_max, y_max))
	for x := 0; x < x_max; x++ {
		for y := 0; y < y_max; y++ {
			c := img.At(x/scale, y/scale)
			new_img.Set(x, y, c)
		}
	}
	return new_img
}

type Coords struct {
	x int
	y int
}

func (c Coords) toString() string {
	return strconv.Itoa(c.x) + "," + strconv.Itoa(c.y)
}

func image2X(img image.Image) image.Image {
	scale := 2
	x_max, y_max := img.Bounds().Max.X*scale, img.Bounds().Max.Y*scale
	new_img := image.NewNRGBA(image.Rect(0, 0, x_max, y_max))
	visited := make(map[string]bool)
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < y_max; y++ {
			c := img.At(x, y)
			new_img.Set(x*scale, y*scale, c)
			coord := Coords{x: x * scale, y: y * scale}
			visited[coord.toString()] = true
		}
	}
	for x := 0; x < x_max; x++ {
		for y := 0; y < y_max; y++ {
			coord := Coords{x, y}
			c := findNearestColor(new_img, coord, scale, visited)
			new_img.Set(x, y, c)
		}
	}
	return new_img
}

func findNearestColor(img image.Image, coord Coords, scale int, visited map[string]bool) color.Color {
	sample := make(map[color.Color]int)
	ratio := scale / 2
	for i := coord.x - ratio; i < coord.x+ratio; i++ {
		for j := coord.y - ratio; j < coord.y+ratio; j++ {
			coord := Coords{i, j}
			v, ok := visited[coord.toString()]
			if ok && v {
				c := img.At(i, j)
				_, ok := sample[c]
				if ok {
					sample[c]++
				} else {
					sample[c] = 1
				}
			}
		}
	}
	var choice color.Color
	count := 0
	for i, v := range sample {
		if count < v {
			choice = i
			count = v
		}
	}
	return choice
}

func main() {
	// img, err := openImage("./Go_gopher.jpg")
	// img, err := openImage("./gopher1.jpg")
	// if err != nil {
	// 	panic(err)
	// }
	r, _ := os.Open("./lil_gopher.png")
	img, _ := png.Decode(r)

	for i := 0; i < 5; i++ {
		img = image2X(img)
	}

	http_image := httpImage{img}
	http.HandleFunc("/", http_image.serveImage)
	http.ListenAndServe(":8080", nil)
}
