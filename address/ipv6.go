package address

import (
  "log"
    "net/http"
    "io/ioutil"
     "strings"
)
func GetIPv6() (string) {
    res, err := http.Get("http://ip6.me/api/")
    if err != nil  {
      log.Println("failed")
    }
    defer res.Body.Close()

      body, _ := ioutil.ReadAll(res.Body)
      ipv6 := strings.Split(string(body), ",")[1]
      return ipv6
}
