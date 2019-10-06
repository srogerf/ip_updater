package provider

import (
    "log"
    "net/http"
     "io/ioutil"
    "time"
    "encoding/json"
    "bytes"
)

//define return and parse body of html
    type PutARecord struct {
         Data  string `json:"data"`
         Port  int `json:"port"`
         Priority  int `json:"priority"`
         Protocol  string `json:"protocol"`
         Service  string `json:"service"`
         Ttl  int `json:"ttl"`
         Weight  int `json:"weight"`
    }

//update the ip address
func UpdateIPv4(domain string, host string, api_key string, api_secret string, value string)  {
    name := "test"
    checkUri := "https://api.godaddy.com/v1/domains/" + domain + "/records/A/" + name
    sso := "sso-key " + api_key + ":" + api_secret

    log.Println("updating %s\n", checkUri)

    body_datas := make([]PutARecord, 1)
    //body_data := new(PutARecord)
    body_datas[0].Data = value
    body_datas[0].Ttl = 600
    body_datas[0].Port = 443

    body_json, err := json.Marshal(body_datas)
    body_json_string := bytes.NewBuffer(body_json)

    body := body_json_string
    if err != nil  {
        log.Println("failed %d\n", err)
    }
    log.Printf("body is %s\n", body)


    req, err := http.NewRequest(http.MethodPut, checkUri, body)
    if err != nil  {
        log.Println("failed %d\n", err)
    }

    req.Header.Set("Authorization", sso)
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Content-Type", "application/json")



        client := &http.Client{Timeout: time.Second * 10}
        res, err := client.Do(req)
        if err != nil  {
            log.Println("failed %d\n", err)
        }
            log.Printf("response %d\n", res)
        defer res.Body.Close()

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
        log.Println("failed %d\n", err)
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
