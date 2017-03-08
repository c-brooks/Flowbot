// main.go is the entry point for the program.
// It sets up all necessary connections for the application.
package main

import (
  "os"
  "github.com/c-Brooks/bADLIB/scraper"
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

  scraper.Scrape(artistName)
}
