package main

import (
    "flag"
    "log"
     "github.com/srogerf/ip_updater/address"
     "github.com/srogerf/ip_updater/provider"
)

func main() {
    log.Print("Starting IP importer")
    api_key := flag.String("key", "none", "godaddy api key")
    api_secret := flag.String("secret", "none", "godaddy api secret")
    domain := flag.String("domain", "none", "domain to update")
    host := flag.String("host", "none", "host to update")
    flag.Parse()

//get current ip values
    assigned_ipv4 := address.GetIPv4()
//    assigned_ipv6 := address.GetIPv6()

//get dns ip values
    dns_ipv4 := provider.GetIPv4(*domain, *host, *api_key, *api_secret)

    log.Printf("\n    assigned IP4: %s\n    dns ip v4: %s\n\n", assigned_ipv4, dns_ipv4)


//check if right
    if assigned_ipv4 != dns_ipv4 {
        log.Println("IPv4 requires update")
    } else {
        log.Println("IPv4 does not require update")
    }

//update if out of date

}
