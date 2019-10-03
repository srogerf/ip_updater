package address

import (
  "log"
    "net/http"
    "io/ioutil"
     "strings"
)
func GetIPv4() (string) {
    res, err := http.Get("http://ip4only.me/api/")
    if err != nil  {
      log.Println("failed")
    }
    defer res.Body.Close()

      body, _ := ioutil.ReadAll(res.Body)
      ipv4 := strings.Split(string(body), ",")[1]
      return ipv4
}
