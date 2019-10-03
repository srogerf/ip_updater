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
    domain := flag.String("doamin", "none", "doamin to update")
    host := flag.String("host", "none", "host to update")

    ipv4 := address.GetIPv4()
    ipv6 := address.GetIPv6()

    log.Printf("\n    IP4: %s\n    IP6: %s\n\n", ipv4, ipv6)

    currentIPv4 := provider.GetIPv4(*domain, *host, *api_key, *api_secret)
    log.Printf("\n    current IP4: %s\n    IP6: \n\n", currentIPv4)
//    provider.godaddy.GetIPv4()
// get current record
//    updater.SelectProvider("godaddy")
//    updater.SelectRecord("test.com")
//    ipv4 = updater.getCurrentRecord()


//check if right


//update if out of date

}
