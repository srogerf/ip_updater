package provider

import (
    "log"
    "net/http"
    "io/ioutil"
    "strings"
)
// test to check api
// curl -X GET -H"Authorization: sso-key ${API_KEY}:${API_SECRET}" "https://api.godaddy.com/v1/domains/${DOMAIN}/records/A/${DOMAIN}"

//d "[ { \"data\": \"$currentIp\", \"port\": $port, \"priority\": 0, \"protocol\": \"string\", \"service\": \"string\", \"ttl\": $ttl, \"weight\": $weight } ]"

func UpdateIPv4() {
}

// get the current ip of the host from godaddy given the domain and host
func GetIPv4(domain string, host string, api_key string, api_secret string) (string) {
    checkUri := "https://api.godaddy.com/v1/domains"
    res, err := http.Get(checkUri)

    if err != nil  {
      log.Println("failed")
    }
    defer res.Body.Close()

      body, _ := ioutil.ReadAll(res.Body)
      ipv4 := strings.Split(string(body), ",")[1]
      return ipv4
}
