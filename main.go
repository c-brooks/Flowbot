// main.go is the entry point for the program.
// It sets up all necessary connections for the application.
package main

import (
  "fmt"
  "log"
  "os"
  "strings"
  "github.com/PuerkitoBio/goquery"
)

func main() {
  var artistName string
  if len(os.Args) > 1 {
    artistName = os.Args[1]
  } else {
    // Falback to a classic
    artistName = "migos"
  }

  // Iterate over tracks
  for _, track := range scrapeTrackList("http://www.azlyrics.com/" + string(artistName[0]) + "/" + artistName + ".html") {
    if track != "" {
      geniusUrl := "https://genius.com/" + artistName + "-" + dasherize(track) + "-lyrics"
      scrapeLyrics(geniusUrl)
    }
  }
}

// Scrape Migos songs from http://www.azlyrics.com/m/migos.html
// Return a list of tracks
func scrapeTrackList(websiteUrl string) []string {
  fmt.Println("GET [", websiteUrl, "]\n")
  doc, err := goquery.NewDocument(websiteUrl)
  if err != nil {
    panic(err.Error())
  }

  var trackList []string
  doc.Find("#listAlbum > a").Each(func (i int, s *goquery.Selection) {
    trackList = append(trackList, s.Text())
  })
  if len(trackList) == 0 {
    log.Fatal("No tracks found!")
  }
  return trackList
}


// Scrape lyrics from Genius
// Print to standard output
func scrapeLyrics(websiteUrl string) {
  fmt.Println("\t GET [", websiteUrl, "]\n")
  doc, err := goquery.NewDocument(websiteUrl)
  if err != nil {
    panic(err.Error())
  }

  fmt.Println(doc.Find(".lyrics").Text())
}

// Change the track name into a url-friendly form
// This includes removing some punctuation for Genius' standard urls
func dasherize(track string) string {
  r := strings.NewReplacer(" ", "-", "(", "", ")", "", "'", "", ".", "", "&", "and")
  return r.Replace(track)
}
