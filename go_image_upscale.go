package main

import (
	"flag"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
)

func main() {
	var input string
	var output string
	var operation string
	var help bool
	flag.StringVar(&input, "input", "", "Input file")
	flag.StringVar(&output, "output", "", "Output file")
	flag.StringVar(&operation, "operations", "", "Operations to perform on the input image")
	flag.BoolVar(&help, "help", false, "Print help")
	flag.Parse()
	if help {
		flag.PrintDefaults()
	}

	r, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	var img image.Image
	i := path.Ext(input)
	i = strings.ToLower(i)
	switch i {
	case ".png":
		img, err = png.Decode(r)
		if err != nil {
			panic(err)
		}
	case ".jpg":
		img, err = jpeg.Decode(r)
		if err != nil {
			panic(err)
		}
	default:
		println("Unknown input type")
		os.Exit(1)
	}

	operations := strings.Split(operation, ",")
	for _, v := range operations {
		switch v {
		case image2x:
			img = image2X(img)
		case scale:
			img = integerScale(img, 2)
		}
	}

	w, err := os.Create(output)
	if err != nil {
		panic(err.Error())
	}
	o := path.Ext(output)
	o = strings.ToLower(o)
	switch o {
	case ".png":
		png.Encode(w, img)
	case ".jpg", ".jpeg":
		jpeg.Encode(w, img, &jpeg.Options{Quality: 100})
	default:
		println("Unknown output type")
		os.Exit(1)
	}
}
