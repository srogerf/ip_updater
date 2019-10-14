/**
https://developer.godaddy.com/doc/endpoint/domains
*/
package provider

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//update the ip address
func UpdateIPv4(domain string, host string, apiKey string, apiSecret string, value string) {
	checkUri := "https://api.godaddy.com/v1/domains/" + domain + "/records/A/" + host
	sso := "sso-key " + apiKey + ":" + apiSecret

	//define the payload structure for the request
	type PutARecord struct {
		Data     string `json:"data"`     //data; req
		Port     int    `json:"port"`     //port is 80, 443; req
		Priority int    `json:"priority"` //10 or inc of 5
		Protocol string `json:"protocol"`
		Service  string `json:"service"`
		Ttl      int    `json:"ttl"`    //600 min; req
		Weight   int    `json:"weight"` //10 or inc of 5
	}

	//create the payload
	bodyData := make([]PutARecord, 1)
	bodyData[0].Data = value
	bodyData[0].Ttl = 600
	bodyData[0].Port = 443
	bodyJson, err := json.Marshal(bodyData)
	if err != nil {
		log.Printf("Failed to create json payload: %s\n", err)
		panic(err)
	}

	body := bytes.NewBuffer(bodyJson)
	req, err := http.NewRequest(http.MethodPut, checkUri, body)
	if err != nil {
		log.Printf("Failed to create update request: %s\n", err)
		panic(err)
	}

	// build the update headers
	req.Header.Set("Authorization", sso)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to update ipv4: %s\n", err)
		panic(err)
	}
	defer res.Body.Close()
}

// get the current ip of the host from godaddy given the domain and host
func GetIPv4(domain string, host string, apiKey string, apiSecret string) string {
	//endpoints and auth
	checkUri := "https://api.godaddy.com/v1/domains/" + domain + "/records/A/" + host
	sso := "sso-key " + apiKey + ":" + apiSecret

	req, err := http.NewRequest("GET", checkUri, nil)
	if err != nil {
		log.Printf("Failed to create request: %s\n", err)
		panic(err)
	}

	//basic header info
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", sso)

	client := &http.Client{Timeout: time.Second * 10}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to send request: %s\n", err)
		panic(err)
	}
	defer res.Body.Close()

	//define return and parse body of html
	type GetARecord struct {
		Data       string `json:"data"`
		Name       string `json:"name"`
		Ttl        int    `json:"ttl"`
		RecordType string `json:"type"`
	}
	var data []GetARecord
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Printf("Failed to read return result: %s\n    Body is %s", err, body)
		panic(err)
	}
	log.Println(data)
	return data[0].Data
}
