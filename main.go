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

func main()  {
  urlArr := GetSongsPath()

  for _, songEndpoint := range urlArr {
    GetSongLyrics(songEndpoint)
  }
  fmt.Println(urlArr)
}

// GetSongsPath
// Returns
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


func scrapeLyrics(websiteUrl string) {
  doc, err := goquery.NewDocument(websiteUrl)
  if err != nil {
    panic(err.Error())
  }
  // Find the 
  fmt.Println(doc.Find(".lyrics").Text())
}


func scrapeLyrics1234(websiteUrl string) {
  // var data interface{}
  res, err := http.Get(websiteUrl)
  if err != nil {
    panic(err.Error())
  }

  fmt.Println(res)
  z := html.NewTokenizer(res.Body)

  for {
    tt := z.Next()
    switch {
    case tt == html.ErrorToken:
      break

    case tt == html.StartTagToken:
      // nop//t := z.Token()

    for _, attr := range z.Token().Attr {
      // fmt.Println(attr)
      if attr.Key == "class" {
        if (attr.Val == "lyrics") {
          z.Next()
          z.Next()
          fmt.Println(z.Text())
        }
      }
    }
  }
}
}

//    switch {
//    case el == html.ErrorToken:
//    	// End of the document, we're done
//        return
//    case el == html.StartTagToken:
//
//     //  t := z.Token()
//
//        default: fmt.Println(z.Token().Val)
//    }
// }
