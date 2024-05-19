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
	x int
	y int
	c color.Color
}

func (c Coords) toString() string {
	return strconv.Itoa(c.x) + "," + strconv.Itoa(c.y)
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
			coord := Coords{x: x * scale, y: y * scale}
			visited[coord.toString()] = true
		}
	}
	for x := 0; x < x_max; x++ {
		for y := 0; y < y_max; y++ {
			coord := Coords{x: x, y: y}
			go func() {
				c := findNearestColor(newImg, coord, scale, visited)
				newImg.Set(c.x, c.y, c.c)
			}()
		}
	}
	return newImg
}

func findNearestColor(img *image.NRGBA, coord Coords, scale int, visited map[string]bool) Coords {
	sample := make(map[color.Color]int)
	ratio := scale / 2
	for i := coord.x - ratio; i < coord.x+ratio; i++ {
		for j := coord.y - ratio; j < coord.y+ratio; j++ {
			coord := Coords{x: i, y: j}
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
	return Coords{x: coord.x, y: coord.y, c: choice}
}
