package provider

import (
    "log"
    "net/http"
     "io/ioutil"
    "time"
    "encoding/json"
)
// test to check api
// curl -X GET -H"Authorization: sso-key ${API_KEY}:${API_SECRET}" "https://api.godaddy.com/v1/domains/${DOMAIN}/records/A/"

//d "[ { \"data\": \"$currentIp\", \"port\": $port, \"priority\": 0, \"protocol\": \"string\", \"service\": \"string\", \"ttl\": $ttl, \"weight\": $weight } ]"

//update the ip address
func UpdateIPv4() {
//    checkUri := "https://api.godaddy.com/v1/domains/" + domain + "/records/A/"
//    sso := "sso-key " + api_key + ":" + api_secret
}

// get the current ip of the host from godaddy given the domain and host
func GetIPv4(domain string, host string, api_key string, api_secret string) (string) {
    checkUri := "https://api.godaddy.com/v1/domains/" + domain + "/records/A/"
    sso := "sso-key " + api_key + ":" + api_secret

    req, err := http.NewRequest("GET", checkUri, nil)
    if err != nil  {
      log.Println("failed")
    }

    req.Header.Set("Accept", "application/json")
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", sso)

    client := &http.Client{Timeout: time.Second * 10}
    res, err := client.Do(req)
    if err != nil  {
      log.Println("failed")
    }
    defer res.Body.Close()

    body, _ := ioutil.ReadAll(res.Body)

//define return and parse body of html
    type ARecord struct {
         Data  string `json:"data"`
         Name  string `json:"name"`
         Ttl  int `json:ttl"`
         Record_type  string `json:type"`
    }
    var data []ARecord
    err = json.Unmarshal([]byte(body), &data);

    if err != nil  {
        log.Println("failed %d\n", err)
    }

    log.Printf("current ip:  %s\n", data[0].Data)
    return data[0].Data
}
