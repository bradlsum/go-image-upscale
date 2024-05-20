package util

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
)

// Read an image from file at the provided path
func ImageFromFile(fpath string) (image.Image, error) {
	r, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}

	var img image.Image
	i := path.Ext(fpath)
	i = strings.ToLower(i)
	switch i {
	case ".png":
		img, err = png.Decode(r)
		if err != nil {
			return nil, err
		}
	case ".jpg":
		img, err = jpeg.Decode(r)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown input type")
	}

	return img, nil
}

// Write the provided image to file
func ImageToFile(img image.Image, fpath string) error {
	w, err := os.Create(fpath)
	if err != nil {
		panic(err.Error())
	}
	o := path.Ext(fpath)
	o = strings.ToLower(o)
	switch o {
	case ".png":
		png.Encode(w, img)
	case ".jpg", ".jpeg":
		jpeg.Encode(w, img, &jpeg.Options{Quality: 100})
	default:
		return fmt.Errorf("unknown output type")
	}
	return nil
}
