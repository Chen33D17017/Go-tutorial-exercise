package main

import ("fmt"
		"strings"
	    "strconv")

type IPAddr [4]byte

type IPtest struct{
	name int
	ip []byte
}


// TODO: Add a "String() string" method to IPAddr.


func (ip IPAddr) String() string {
	
	ips := make([]string, 0, len(ip))
	for _, i := range ip{
		ips = append(ips, strconv.Itoa(int(i)))
	}
	return strings.Join(ips, ".")
	
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	
	
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
