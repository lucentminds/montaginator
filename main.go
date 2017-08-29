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
	"os"
	"io/ioutil"
	"image"
	"image/color"
	"log"
)

import "github.com/disintegration/imaging"

// Required boilerplate entry function for all go applications.
func main() {
	// Start defining your application here.
	cPathImages := os.Args[1]
	cOutputFile := os.Args[2]
	fmt.Println( cPathImages )
	fmt.Println( cOutputFile )

	// Scan the directory for the image files.
	aFiles, _ := ioutil.ReadDir( cPathImages )
	nWidthOne, nHeightOne := getImageDimension( cPathImages+"/"+aFiles[0].Name() )
	nHeightAll := nHeightOne * len( aFiles )


	imgMontage := imaging.New( nWidthOne, nHeightAll, color.NRGBA{0, 0, 0, 0})
	nY := 0

	for _, oFile := range aFiles{
		//fmt.Println( oFile.Name() )

		// Open the next image.
		src, err := imaging.Open( cPathImages+"/"+oFile.Name() )
		if err != nil {
			log.Fatalf("Open failed: %v", err)
		}

		imgMontage = imaging.Paste( imgMontage, src, image.Pt( 0, nHeightOne*nY ) )
		nY++
	}// /for()

	imaging.Save( imgMontage, cOutputFile )
}// /main()

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
}