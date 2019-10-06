package main

import (
	"flag"
	"github.com/srogerf/ip_updater/address"
	"github.com/srogerf/ip_updater/provider"
	"log"
)

func main() {
	log.Print("Starting IP updater")
	apiKey := flag.String("key", "none", "godaddy api key")
	apiSecret := flag.String("secret", "none", "godaddy api secret")
	domain := flag.String("domain", "none", "domain to update")
	host := flag.String("host", "@", "host to update")
	flag.Parse()

	//get current ip values
	assignedIpv4 := address.GetIPv4()
	assignedIpv6 := address.GetIPv6()
	log.Printf("\n    assigned IP4: %s\n    assigned IP6: %s\n", assignedIpv4, assignedIpv6)

	//get dns ip values
	dnsIpv4 := provider.GetIPv4(*domain, *host, *apiKey, *apiSecret)
	log.Printf("\n    dns IP4: %s\n", dnsIpv4)

	//update the ipv4 record
	if assignedIpv4 != dnsIpv4 {
		log.Println("IPv4 requires update")
		provider.UpdateIPv4(*domain, *host, *apiKey, *apiSecret, assignedIpv4)
	} else {
		log.Println("IPv4 does not require update")
	}
}
