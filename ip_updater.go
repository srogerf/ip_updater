package main

import (
    "flag"
    "log"
     "github.com/srogerf/ip_updater/address"
     "github.com/srogerf/ip_updater/provider"
)

func main() {
    log.Print("Starting IP updater")
    api_key := flag.String("key", "none", "godaddy api key")
    api_secret := flag.String("secret", "none", "godaddy api secret")
    domain := flag.String("domain", "none", "domain to update")
    host := flag.String("host", "@", "host to update")
    flag.Parse()

//get current ip values
    assigned_ipv4 := address.GetIPv4()
    assigned_ipv6 := address.GetIPv6()
    log.Printf("\n    assigned IP4: %s\n    assigned IP6: %s\n", assigned_ipv4, assigned_ipv6)

//get dns ip values
    dns_ipv4 := provider.GetIPv4(*domain, *host, *api_key, *api_secret)
    log.Printf("\n    dns IP4: %s\n", dns_ipv4)

//update the ipv4 record
    if assigned_ipv4 != dns_ipv4 {
        log.Println("IPv4 requires update")
        provider.UpdateIPv4(*domain, *host, *api_key, *api_secret, assigned_ipv4)
    } else {
        log.Println("IPv4 does not require update")
    }
}
