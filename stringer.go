package main

import "fmt"

type IPAddr [6]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	var ipa string
	for _, val := range ip {
		ipa += fmt.Sprintf("%v.", val)
	}
	return ipa[:len(ipa)-1]
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1, 9, 8},
		"googleDNS": {8, 8, 8, 8, 4, 4},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
