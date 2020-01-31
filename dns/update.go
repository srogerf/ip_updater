package dns

import (
	"github.com/srogerf/ip_updater/address"
	"github.com/srogerf/ip_updater/provider"
	"log"
)

func Update(domain string, host string, apiKey string, apiSecret string) {
	//get current ip values
	assignedIpv4 := address.GetIPv4()
	assignedIpv6 := address.GetIPv6()
	log.Printf("\n    current IP4: %s\n    current IP6: %s\n", assignedIpv4, assignedIpv6)

	//get dns ip values
	dnsIpv4 := provider.GetIPv4(domain, host, apiKey, apiSecret)
	log.Printf("\n    dns IP4: %s\n", dnsIpv4)

	//update the ipv3 record
	if assignedIpv4 != dnsIpv4 {
		log.Println("IPv4 requires update")
		provider.UpdateIPv4(domain, host, apiKey, apiSecret, assignedIpv4)
	} else {
		log.Println("IPv4 does not require update")
	}
}
