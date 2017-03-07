package main

import (
	"io/ioutil"
	"log"
	"strings"
	"os"
	"fmt"
)

func main() {
        // Set the count to 0
	count := 0
        // Read current directory, ReadDir also returns an error
	files, err := ioutil.ReadDir(".")
        // Handle the error
	if err != nil { log.Fatal(err) }
        // Iterate through the files in the directory
	for _, file := range files {
                // In my case, if the file didn't contain a dot, continue
		if !strings.Contains(file.Name(), ".") {continue}
                // If it was Gogland's default .idea folder, continue
		if file.Name() == ".idea" {continue}
                // If the file was a .go file, continue
		if strings.HasSuffix(file.Name(), ".go") {continue}
                // Rename the file to the current name, but remove the '.'
                // Here you can customise it to something else
		os.Rename(file.Name(), strings.Replace(file.Name(), ".", "", 1))
                // When all that's done, increment our count
		count++
	}
        // Print the outcome
	fmt.Println("All done. Changed", count, "files")
}
