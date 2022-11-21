package filemanager

import (
	"io"
	"log"
	"os"
	"strings"
)

// Gets the name of the last image that was generated and returns a string of the filename.
// If a file doesn't exist then start at 1
func GetCurrentImageName() string {
	fileName := strings.Split(getCurrentFileName(), ".")[0]
	if fileName == "" {
		return "0"
	}
	// @todo will need to get last file from old folder to get latest integer
	return fileName
}

// Internal function for getting the current full file name
func getCurrentFileName() string {
	files, err := os.ReadDir("./current")
	if err != nil {
		log.Fatal(err)
	}

	fileName := ""
	if len(files) == 1 {
		file := files[0]
		fileName = file.Name()
	}
	return fileName
}

// Moves the current image to the old folder
func MoveCurrentImage() {
	fileName := getCurrentFileName()
	if fileName == "" {
		return
	}
	currentFile, err := os.Open("./current/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	oldFile, err := os.Create("./old/" + fileName)
	if err != nil {
		currentFile.Close()
		log.Fatal(err)
	}
	defer oldFile.Close()
	// Files are in buffer and time to copy
	_, err = io.Copy(oldFile, currentFile)
	currentFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	// Remove the current file
	err = os.Remove("./current/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
}
