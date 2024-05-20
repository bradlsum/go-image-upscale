package operations

import (
	"image"
	"image/color"
	"strconv"
)

const (
	Image2x = "image2x"
	Scale   = "scale2x"
)

var Operations = []string{Image2x, Scale}

func IntegerScale(img image.Image, scale int) *image.NRGBA {
	x_max, y_max := img.Bounds().Max.X*scale, img.Bounds().Max.Y*scale
	newImg := image.NewNRGBA(image.Rect(0, 0, x_max, y_max))
	for x := 0; x < x_max; x++ {
		for y := 0; y < y_max; y++ {
			c := img.At(x/scale, y/scale)
			newImg.Set(x, y, c)
		}
	}
	return newImg
}

type Coords struct {
	X int
	Y int
	C color.Color
}

func (c Coords) toString() string {
	return strconv.Itoa(c.X) + "," + strconv.Itoa(c.Y)
}

func Image2X(img image.Image) *image.NRGBA {
	scale := 2
	x_max, y_max := img.Bounds().Max.X*scale, img.Bounds().Max.Y*scale
	newImg := image.NewNRGBA(image.Rect(0, 0, x_max, y_max))
	visited := make(map[string]bool)
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < y_max; y++ {
			c := img.At(x, y)
			newImg.Set(x*scale, y*scale, c)
			coord := Coords{X: x * scale, Y: y * scale}
			visited[coord.toString()] = true
		}
	}
	for x := 0; x < x_max; x++ {
		for y := 0; y < y_max; y++ {
			coord := Coords{X: x, Y: y}
			go func() {
				c := FindNearestColor(newImg, coord, scale, visited)
				newImg.Set(c.X, c.Y, c.C)
			}()
		}
	}
	return newImg
}

func FindNearestColor(img *image.NRGBA, coord Coords, scale int, visited map[string]bool) Coords {
	sample := make(map[color.Color]int)
	ratio := scale / 2
	for i := coord.X - ratio; i < coord.X+ratio; i++ {
		for j := coord.Y - ratio; j < coord.Y+ratio; j++ {
			coord := Coords{X: i, Y: j}
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
	return Coords{X: coord.X, Y: coord.Y, C: choice}
}
