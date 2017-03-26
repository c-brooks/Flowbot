// main.go is the entry point for the program.
package main

import (
	"os"
	"github.com/c-Brooks/bADLIB/scraper"
	"fmt"
	"github.com/c-Brooks/bADLIB/ml"
)

func main() {
	// Get artist name from command-line args
	var artistName string
	if len(os.Args) > 1 {
		artistName = os.Args[1]
	} else {
		// Falback to a classic
		artistName = "migos"
	}

	fmt.Println(artistName)
	songArr := scraper.Scrape(artistName)
	for song := range songArr {
		ml.Train(songArr[song], 1)
	}
}
