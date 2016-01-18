package main

import (
	"fmt"
	"image"
	"log"

	"github.com/disintegration/imaging"
	"github.com/jteeuwen/go-pkg-optarg"
)

func main() {
	optarg.Header("General options")
	optarg.Add("h", "height", "image height", 100)
	optarg.Add("w", "width", "image width", 0)
	optarg.Add("i", "input", "input filename", "")
	optarg.Add("o", "output", "output filename", "")

	var (
		height  int
		width   int
		inName  string
		outName string
	)

	for opt := range optarg.Parse() {
		switch opt.ShortName {
		case "h":
			height = opt.Int()
		case "w":
			width = opt.Int()
		case "i":
			inName = opt.String()
		case "o":
			outName = opt.String()
		}
	}

	if inName == "" || outName == "" {
		optarg.Usage()
		log.Fatal()
	}

	orgImg, err := imaging.Open(inName)
	if err != nil {
		log.Fatal(err)
	}

	dstImg := resize(orgImg, width, height)

	err = imaging.Save(dstImg, outName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Resize to", outName)
}

func resize(img image.Image, width int, height int) image.Image {
	dstImg := imaging.Resize(img, width, height, imaging.Lanczos)
	return dstImg.SubImage(dstImg.Rect)
}
