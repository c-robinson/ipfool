package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/c-robinson/iplib"
)

type addrType int

var (
	v46 addrType = 0
	v4  addrType = 4
	v6  addrType = 6
)

func retrieveIPAddress(s string, t addrType) net.IP {
	ip := net.ParseIP(s)

	if ip == nil {
		fmt.Println("supplied value is not a valid IP address")
		os.Exit(1)
	}

	if t != addrType(0) {
		if int(t) != iplib.EffectiveVersion(ip) {
			fmt.Println("supplied IP version not supported in this function")
			os.Exit(1)
		}
	}

	return ip
}

func retrieveIPNetwork(s string, t addrType) iplib.Net {
	_, ipnet, err := iplib.ParseCIDR(s)
	if err != nil {
		fmt.Println("supplied value is not a valid IP netblock")
		os.Exit(1)
	}

	return ipnet
}
