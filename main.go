// main.go is the entry point for the program.
// It sets up all necessary connections for the application.
package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "strings"
)

func main() {

  // Iterate over tracks
  for _, track := range scrapeTrackList("http://www.azlyrics.com/m/migos.html") {
    if track != "" {
      geniusUrl := "https://genius.com/Migos-" + dasherize(track) + "-lyrics"
      scrapeLyrics(geniusUrl)
    }
  }
}

// Scrape Migos songs from http://www.azlyrics.com/m/migos.html
// Return a list of tracks
func scrapeTrackList(websiteUrl string) []string {
  doc, err := goquery.NewDocument(websiteUrl)
  if err != nil {
    panic(err.Error())
  }

  var trackList []string
  doc.Find("#listAlbum > a").Each(func (i int, s *goquery.Selection) {
    trackList = append(trackList, s.Text())
  })
  return trackList
}


// Scrape lyrics from Genius
// Print to standard output
func scrapeLyrics(websiteUrl string) {
  fmt.Println("\n === Scraping from", websiteUrl, "\n")
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
