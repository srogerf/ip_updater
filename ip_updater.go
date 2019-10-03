package main

import (
    "log"
    "net/http"
    "io/ioutil"
     "strings"
)

func main() {
  log.Print("Starting IP importer")
  res, err := http.Get("http://ip4only.me/api/")
  if err != nil  {
    log.Println("failed")
  }
  defer res.Body.Close()

    body, _ := ioutil.ReadAll(res.Body)
    ipv4 := strings.Split(string(body), ",")[1]
    log.Printf("Got ipv4 %s\n", ipv4)
//    bodySlice  := strings.Split(body, ",")
//    log.Printf("Got %d\n", bodySlice)
}
