package main

// Our imports
import (
	"io/ioutil"
	"log"
	"strings"
	"os"
	"fmt"
)

// Our main method
func main() {
	arrangeScreenShots()
}

// Creating our function, I like creating functions rather than doing it all in main. Sue me
func arrangeScreenShots() {
	// Getting our possible screenshots/images in this example, there are loads more, I know, you can add them
	var imageExtensionList = []string{"jpeg","jpg","png", "gif"}
	// Creating our variables for the files and the error in the '.' directory, which will always be the current directory
	files, err := ioutil.ReadDir(".")
	// Check if there are any errors, if so, let us know
	if err != nil { log.Fatal(err) }
	// Iterate through the files
	for _, file := range files {
		// Checks if the files end with .png, .jpg, .jpeg
		for _, extensions := range imageExtensionList {
			// Iterate through the list, idk why I didn't do this before lol
			if strings.HasSuffix(file.Name(), extensions) {
				// This method moves the files into the new directory, 'Screenshots' and gives them their same name.
				// You have to give it it's own name while moving, just how it works.
				os.Rename(file.Name(), "Screenshots/" + file.Name())
			}
		}
	}
	// Let us know when the process is done.
	fmt.Println("Moved all the screenshots.")
}
