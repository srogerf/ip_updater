package address

import (
  "log"
    "net/http"
    "io/ioutil"
     "strings"
)

func GetIPv6() (string) {
    checkUri := "http://ip6.me/api/"

    res, err := http.Get(checkUri)
    if err != nil  {
        log.Printf("Failed to send ip6 discovery request: %s\n", err)
        panic(err)
    }
    defer res.Body.Close()

    body, _ := ioutil.ReadAll(res.Body)
    ipv6 := strings.Split(string(body), ",")[1]
    return ipv6
}
