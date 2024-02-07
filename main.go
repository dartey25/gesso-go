package main

import (
	"fmt"
	"os"

	"github.com/h2non/bimg"
)

func main() {
	options := bimg.Options{
		Width:     1920,
		Height:    1080,
		Crop:      true,
		Quality:   95,
		Interlace: true,
	}

	buffer, err := bimg.Read("assets/img3.jpg")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Process(options)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("new.avif", newImage)
}
