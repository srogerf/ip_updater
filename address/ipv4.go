package address

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/**
  Get the ipv4 address using a free responder
*/
func GetIPv4() string {
	checkUri := "http://ip4only.me/api/"
	res, err := http.Get(checkUri)

	if err != nil {
		log.Printf("Failed to send ip4 discovery request: %s\n", err)
		panic(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	ipv4 := strings.Split(string(body), ",")[1]
	return ipv4
}
