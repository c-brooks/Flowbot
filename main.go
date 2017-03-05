// main.go is the entry point for the program.
// It sets up all necessary connections for the application.
package main

import (
  "log"
  "net/http"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "os"
)

func main()  {
  urlArr := GetSongsPath()
  fmt.Println(urlArr)
}

func GetSongsPath() []string {
  var data interface{}

  accessToken := os.Getenv("GENIUS_ACCESS_TOKEN")
  migosId := "44080" // Genius artist id for Migos

  artistEndpoint := "https://api.genius.com/artists/" + migos_id + "/songs?access_token=" + accessToken
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
     ret = append(ret, "https://api.genius.com/" + song.(map[string]interface{})["api_path"].(string))
   }
   return ret
}
