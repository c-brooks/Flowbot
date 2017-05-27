// Package main is the entry point for the program.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-Brooks/Flowbot/ml"
	"github.com/c-Brooks/Flowbot/scraper"
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
	songArr := scraper.GetByArtist(artistName)
	fmt.Println(songArr)
	songs := strings.Join(songArr, "\n")
	ml.Train(songs, 3)
}
