package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/bradlsum/gscale/internal/operations"
	"github.com/bradlsum/gscale/internal/util"
)

func main() {
	// Setup command-line arguments
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

	// Open the input file
	img, err := util.ImageFromFile(input)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	// Run the requested operations of the image
	operationsSlice := strings.Split(operation, ",")
	for _, v := range operationsSlice {
		switch v {
		case operations.Image2x:
			img = operations.Image2X(img)
		case operations.Scale:
			img = operations.IntegerScale(img, 2)
		}
	}

	// Write the processed image to file
	err = util.ImageToFile(img, output)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
