package scraper


import (
  "fmt"
  "log"
  "strings"
  "github.com/PuerkitoBio/goquery"
)

func Scrape(artistName string)  {
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
    log.Fatal("No tracks found!") // exit
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

  // Put the lyrics into a table
  // Returns a 2-D array of lines, words
  formatLyrics(doc.Find(".lyrics").Text())
}

// Change the track name into a url-friendly form
// This includes removing some punctuation for Genius' standard urls
func dasherize(track string) string {
  r := strings.NewReplacer(" ", "-", "(", "", ")", "", "'", "", ".", "", "&", "and")
  return r.Replace(track)
}


// Format lyrics into a 2-D array
// Returns Array.<Array.<string>>
func formatLyrics(lyrics string) {
  var lyricsArr [][]string
  fmt.Println(lyrics)

  for _, line := range strings.Split(lyrics, "\n") {
    // Test for unwanted lines
    line = strings.Trim(line, " ")
    if len(line) > 0 && string(line[0]) != "[" {
      tempRow := strings.Split(line, " ")
      lyricsArr = append(lyricsArr, tempRow)
    }
  }
  fmt.Println(lyricsArr)
}
