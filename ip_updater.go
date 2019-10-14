package main

import (
	"flag"
	"github.com/srogerf/ip_updater/dns"
	"log"
)

func main() {
	log.Print("Starting IP updater")
	apiKey := flag.String("key", "none", "godaddy api key")
	apiSecret := flag.String("secret", "none", "godaddy api secret")
	domain := flag.String("domain", "none", "domain to update")
	host := flag.String("host", "@", "host to update")
	server := flag.Bool("daemon", true, "Start as server")
	flag.Parse()

	if *server {
		log.Println("starting a server")
	}
	dns.Update(*domain, *host, *apiKey, *apiSecret)
}
