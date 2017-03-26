package scraper


import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func Scrape(artistName string) []string {
	var songBuf bytes.Buffer
	var retArr []string
	for _, track := range scrapeTrackList("http://www.azlyrics.com/" + string(artistName[0]) + "/" + artistName + ".html") {
    	if track != "" {
			geniusUrl := "https://genius.com/" + artistName + "-" + dasherize(track) + "-lyrics"
			songBuf.WriteString(scrapeLyrics(geniusUrl))
			retArr = append(retArr, songBuf.String())
			break
    }
  }
	fmt.Println(retArr)
	return retArr
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
func scrapeLyrics(websiteUrl string) string {
	fmt.Println("\t GET [", websiteUrl, "]\n")
	doc, err := goquery.NewDocument(websiteUrl)
	if err != nil {
		panic(err.Error())
	}
	return formatLyrics(doc.Find(".lyrics").Text())
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
		if len(lowerLine) > 0 && string(lowerLine[0]) != "[" {
			retLyrics.WriteString(lowerLine + " ")
		}
	}
	return retLyrics.String()
}
