// Package scraper contains routines to scrape song data
package scraper

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetByArtist takes an artist name
// and returns an array of songs, (cleaned), as strings.
func GetByArtist(artistName string) []string {
	var retArr []string
	successReqs, failReqs := 0, 0
	resc, errc := make(chan string), make(chan error)

	songList := scrapeTrackList("http://www.azlyrics.com/" + string(artistName[0]) + "/" + artistName + ".html")
	songList = filterEmpty(songList)

	songLen := len(songList)

	for _, track := range songList {
		go func(track string) {
			geniusURL := "https://genius.com/" + artistName + "-" + dasherize(track) + "-lyrics"

			lyrics, err := scrapeLyrics(geniusURL)

			// Handle err
			if err != nil {
				songLen--
				failReqs++
				errc <- err
				return
			}

			// Handle empty response
			if lyrics == "" {
				songLen--
				return
			}

			// Successful
			successReqs++
			resc <- lyrics
		}(track)
	}

	// Build return array from data in channels
	for i := 0; i < songLen; i++ {
		select {
		case res := <-resc:
			retArr = append(retArr, res)
		case err := <-errc:
			log.Println(err)
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
		log.Fatal("No tracks found!") // exit
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

// Change the track name into a url-friendly form.
// This includes removing some punctuation for Genius' standard urls.
func dasherize(track string) string {
	r := strings.NewReplacer(" ", "-", "(", "", ")", "", "'", "", ".", "", "&", "and")
	return r.Replace(track)
}

// Formats lyrics into a cleaned-up string.
func formatLyrics(lyrics string) string {
	var retLyrics bytes.Buffer

	for _, line := range strings.Split(lyrics, "\n") {
		lowCaseLine := strings.ToLower(line)
		lowCaseLine = strings.Trim(lowCaseLine, " ")

		// Test for unwanted lines
		if len(lowCaseLine) > 0 && string(lowCaseLine[0]) != "[" && string(lowCaseLine[0]) != "(" {
			// Separate commas into their own words
			lowCaseLine = strings.NewReplacer(",", " ,").Replace(lowCaseLine)
			retLyrics.WriteString(lowCaseLine + " ")
		}
	}
	return retLyrics.String()
}

// Filters an array of strings, rejecting empty strings
func filterEmpty(arr []string) []string {
	var ret []string

	for _, elem := range arr {
		if elem != "" {
			ret = append(ret, elem)
		}
	}
	return ret
}
