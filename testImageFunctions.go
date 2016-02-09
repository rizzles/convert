package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"path/filepath"
	"os"

	"log"

	"github.com/rwcarlsen/goexif/exif"
)


func main() {
	output, err := ConvertToJpeg("test.png")
	if err == nil {
		extension := filepath.Ext(output)
		name := output[0:len(output) - len(extension)]
		newName := fmt.Sprintf("%v.jpg", name)
		os.Rename(output, newName)

		fmt.Println(newName)
		FixOrientation(newName)
	} else {
		log.Fatal(err)
	}

}

func FixOrientation(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal("error decoding image")
		log.Fatal(err)
	}

	camModel, _ := x.Get(exif.Model) // normally, don't ignore errors!
	fmt.Println(camModel.StringVal())

	return "", nil
}

func ConvertToJpeg(filename string) (string, error) {
	pngImageFile, err := os.Open(filename)

	if err != nil {
		return "", err
	}

	defer pngImageFile.Close()

	pngSource, err := png.Decode(pngImageFile)

	if err != nil {
		return "", err
	}

	jpegImage := image.NewRGBA(pngSource.Bounds())

	draw.Draw(jpegImage, jpegImage.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	draw.Draw(jpegImage, jpegImage.Bounds(), pngSource, pngSource.Bounds().Min, draw.Over)

	outfile := fmt.Sprintf("%s_jpg", filename)
	jpegImageFile, err := os.Create(outfile)

	if err != nil {
		return "", err
	}

	defer jpegImageFile.Close()

	var options jpeg.Options
	options.Quality = 100

	err = jpeg.Encode(jpegImageFile, jpegImage, &options)

	if err != nil {
		fmt.Printf("JPEG Encoding Error: %v\n", err)
		os.Exit(1)
	}

	return outfile, nil
}