package main

import (
	"filemanager"
	"imagebuilder"
)

// Fun image loop
func main() {

	// get image and return number of file
	currentFileName := filemanager.GetCurrentImageName()

	//move image
	filemanager.MoveCurrentImage()

	// do iteration and build new current image
	imagebuilder.BuildImage(currentFileName)
}
