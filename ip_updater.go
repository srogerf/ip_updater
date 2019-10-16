package main

import (
	"flag"
	"github.com/srogerf/ip_updater/dns"
	"github.com/srogerf/ip_updater/server"
	"log"
)

func main() {
	log.Print("Starting IP updater")
	apiKey := flag.String("key", "none", "godaddy api key")
	apiSecret := flag.String("secret", "none", "godaddy api secret")
	domain := flag.String("domain", "none", "domain to update")
	host := flag.String("host", "@", "host to update")
	daemonize := flag.Bool("daemon", false, "Start as server")
	flag.Parse()

	if *daemonize {
		server.Runner()
	} else {
		dns.Update(*domain, *host, *apiKey, *apiSecret)
	}
}
