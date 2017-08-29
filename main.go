/**
 * 08-28-2017
 * Assembles one or more images together..
 * Generated using the go-app template.
 * ~~ Scott Johnson
 */

// Required boilerplate package for all go applications.
package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"os"
)

import "github.com/disintegration/imaging"

// Required boilerplate entry function for all go applications.
func main() {
	// Start defining your application here.
	//fmt.Println( os.Args )

	// Determines the path to the images directory.
	cPathImages := os.Args[1]

	// Determines the path to the output image file.
	cOutputFile := os.Args[2]

	// Verify images directory
	lExists, err := exists( cPathImages )

	if err != nil {
		// Failed to verify images directory.
		log.Fatalf("exists failed: %v", err)
	}

	if !lExists {
		// Images directory does not exist.
		log.Fatalf("error: Directory \"%v\" does not exist.", cPathImages)
	}

	// Scan the directory for the image files.
	aFiles, _ := ioutil.ReadDir(cPathImages)

	// Determine the width and height of the first image.
	// All other images should be the same.
	nWidthOne, nHeightOne := getImageDimension(cPathImages + "/" + aFiles[0].Name())

	// Determines the total height the final montage image will be.
	nHeightAll := nHeightOne * len(aFiles)

	// Determines the main montage image object.
	imgMontage := imaging.New(nWidthOne, nHeightAll, color.NRGBA{0, 0, 0, 0})
	nY := 0

	// Loop over each image in the images directory.
	for _, oFile := range aFiles {

		// Open the next image.
		src, err := imaging.Open(cPathImages + "/" + oFile.Name())

		if err != nil {
			// Failed to open the next image.
			log.Fatalf("Open failed: %v", err)
		}

		// Paste the next image into the main montage image below the previous
		// image.
		imgMontage = imaging.Paste(imgMontage, src, image.Pt(0, nHeightOne*nY))
		nY++
	} // /for()

	// Save the final montage image file.
	imaging.Save(imgMontage, cOutputFile)
} // /main()

func getImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
} // /getImageDimension()

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}