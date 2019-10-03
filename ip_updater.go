package main

import (
    "log"
     "github.com/srogerf/ip_updater/address"
)

func main() {
    log.Print("Starting IP importer")
    ipv4 := address.GetIPv4()
    ipv6 := address.GetIPv6()


    log.Printf("\n    IP4: %s\n    IP6: %s\n\n", ipv4, ipv6)
}
