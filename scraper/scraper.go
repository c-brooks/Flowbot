// Package scraper contains routines to scrape song data
package scraper

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// Scrape is the entrypoint for the scraping routine
// It takes an artist name and returns an array of songs, (cleaned), as strings.
func Scrape(artistName string) []string {
	var retArr []string
	// var wg sync.WaitGroup

	successReqs, failReqs := 0, 0

	resc, errc := make(chan string), make(chan error)

	songList := scrapeTrackList("http://www.azlyrics.com/" + string(artistName[0]) + "/" + artistName + ".html")

	fmt.Println(len(songList))
	for _, track := range songList {

		go func(track string) {
			geniusURL := "https://genius.com/" + artistName + "-" + dasherize(track) + "-lyrics"
			lyrics, err := scrapeLyrics(geniusURL)
			if err != nil {
				failReqs++
				errc <- err
				return
			}
			successReqs++
			resc <- lyrics
		}(track)
	}

	for i := 0; i < len(songList); i++ {
		select {
		case res := <-resc:
			retArr = append(retArr, res)
		case err := <-errc:
			fmt.Println(err)
		}
	}

	fmt.Println("Done.")
	fmt.Println("SUCCESS:  . . . . . . . . .", successReqs)
	fmt.Println("FAIL: . . . . . . . . . . .", failReqs)
	return retArr
}

// Scrape Migos songs from http://www.azlyrics.com/m/migos.html
// Return a list of tracks
func scrapeTrackList(websiteURL string) []string {
	fmt.Println("GET [", websiteURL, "]")
	doc, err := goquery.NewDocument(websiteURL)
	if err != nil {
		log.Println(err.Error())
	}

	var trackList []string
	doc.Find("#listAlbum > a").Each(func(i int, s *goquery.Selection) {
		trackList = append(trackList, s.Text())
	})
	if len(trackList) == 0 {
		log.Println("No tracks found!") // exit
	}
	return trackList
}

// Scrape lyrics from Genius
// Print to standard output
func scrapeLyrics(websiteURL string) (string, error) {
	fmt.Println("\t GET [", websiteURL, "]")
	doc, err := goquery.NewDocument(websiteURL)
	if err != nil {
		return "", err
	}
	return formatLyrics(doc.Find(".lyrics").Text()), err
}

// Change the track name into a url-friendly form
// This includes removing some punctuation for Genius' standard urls
func dasherize(track string) string {
	r := strings.NewReplacer(" ", "-", "(", "", ")", "", "'", "", ".", "", "&", "and")
	return r.Replace(track)
}

// Formats lyrics into a cleaned-up string
// Returns string
func formatLyrics(lyrics string) string {
	var retLyrics bytes.Buffer

	for _, line := range strings.Split(lyrics, "\n") {
		lowerLine := strings.ToLower(line)
		// Test for unwanted lines
		lowerLine = strings.Trim(lowerLine, " ")
		if len(lowerLine) > 0 && string(lowerLine[0]) != "[" && string(lowerLine[0]) != "(" {
			// Separate commas into their own words
			lowerLine = strings.NewReplacer(",", " ,").Replace(lowerLine)
			retLyrics.WriteString(lowerLine + " ")
		}
	}
	return retLyrics.String()
}

////////////////////////////////////////////////////////
///////////////////// NOT USED /////////////////////////
////////////////////////////////////////////////////////

// ConcurrentSlice is a concurrency-safe slice
// to be shared between goroutines
type ConcurrentSlice struct {
	sync.RWMutex
	items []string
}

// ConcurrentSliceItem is an element of a ConcurrentSlice
type ConcurrentSliceItem struct {
	Index int
	Value interface{}
}

// Append appends an item to the concurrent slice
func (cs *ConcurrentSlice) Append(item string) {
	cs.Lock()
	defer cs.Unlock()

	cs.items = append(cs.items, item)
}
