// main.go is the entry point for the program.
// It sets up all necessary connections for the application.
package main

import (
  "net/http"
  "golang.org/x/net/html"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "os"
  "github.com/PuerkitoBio/goquery"
)

func main() {
  urlArr := GetSongsPath()

  for _, songEndpoint := range urlArr {
    GetSongLyrics(songEndpoint)
  }
  fmt.Println(urlArr)
}

// GetSongsPath
// Returns an array of URLs for Migos songs
func GetSongsPath() []string {
  var data interface{}

  accessToken := os.Getenv("GENIUS_ACCESS_TOKEN")
  migosId := "44080" // Genius artist id for Migos

  artistEndpoint := "https://api.genius.com/artists/" + migosId + "/songs?access_token=" + accessToken
  fmt.Println("GET:", artistEndpoint)
  res, err := http.Get(artistEndpoint)
  if err != nil {
    panic(err.Error())
  }

  body, err := ioutil.ReadAll(res.Body)
   if err != nil {
     panic(err.Error())
   }
   err = json.Unmarshal(body, &data)
   if err != nil {
     panic(err.Error())
   }

  songs := data.(map[string]interface{})["response"].(map[string]interface{})["songs"]
  var ret []string
   for _, song := range songs.([]interface{}) {
     ret = append(ret, "https://api.genius.com" + song.(map[string]interface{})["api_path"].(string))
   }
   return ret
}

// Since Genius API is whack, we need to scrape the data from the website
func GetSongLyrics(apiPath string)  {

  var data interface{}
  accessToken := os.Getenv("GENIUS_ACCESS_TOKEN")
  authEndpoint := apiPath + "?access_token=" + accessToken
  fmt.Println("GET:", authEndpoint)
  res, err := http.Get(authEndpoint)
  if err != nil {
    panic(err.Error())
  }

  body, err := ioutil.ReadAll(res.Body)
   if err != nil {
     panic(err.Error())
   }
   err = json.Unmarshal(body, &data)
   if err != nil {
     panic(err.Error())
   }

   song := data.(map[string]interface{})["response"].(map[string]interface{})["song"]
   websiteUrl := "https://genius.com" + song.(map[string]interface{})["path"].(string)
   fmt.Println(" === SCRAPING", websiteUrl, "===")
   scrapeLyrics(websiteUrl)
}

// Scrape lyrics using goQuery
func scrapeLyrics(websiteUrl string) {
  doc, err := goquery.NewDocument(websiteUrl)
  if err != nil {
    panic(err.Error())
  }

  fmt.Println(doc.Find(".lyrics").Text())
}
