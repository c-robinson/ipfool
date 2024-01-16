package cmd

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"

	"github.com/c-robinson/iplib/v2"
)

type addrType int

var (
	v46 addrType = 0
	v4  addrType = 4
	v6  addrType = 6
)

func respondToTrueFalseQuestion(result, exitCodeOnly bool) {
	if !exitCodeOnly {
		fmt.Printf("%t\n", result)
	}
	if result {
		os.Exit(0)
	}
	os.Exit(2)
}

func retrieveIPAddress(s string, t addrType) net.IP {
	ip := net.ParseIP(s)

	if ip == nil {
		fmt.Println("supplied value is not a valid IP address")
		os.Exit(1)
	}

	if t != addrType(0) {
		if int(t) != iplib.EffectiveVersion(ip) {
			fmt.Printf("supplied IP '%s' invalid version for this function\n", ip)
			os.Exit(1)
		}
	}

	return ip
}

func retrieveIPNetwork(s string, t addrType) iplib.Net {
	ipnet := retrieveIPNetwork6WithHostmask(s)
	if ipnet.IP() != nil {
		return ipnet
	}

	_, ipnet, err := iplib.ParseCIDR(s)
	if err != nil {
		fmt.Println("supplied value is not a valid IP netblock")
		os.Exit(1)
	}

	return ipnet
}

func retrieveIPNetwork6WithHostmask(s string) iplib.Net {
	// if there's a hostmask this regex will match it, returning
	// 0: the whole string	2001:db8::/64:24
	// 1: the IP address	2001:db8::
	// 2: the netmask		64
	// 3: the hostmask		24
	r := regexp.MustCompile(`([A-Fa-f0-9:].*)\/([0-9].*):([0-9]{0,3})`)

	ipregex := r.FindStringSubmatch(s)
	if len(ipregex) != 4 {
		return iplib.Net6{}
	}
	address := net.ParseIP(ipregex[1])
	netmask, err := strconv.Atoi(ipregex[2])
	if err != nil {
		return iplib.Net6{}
	}
	hostmask, err := strconv.Atoi(ipregex[3])
	if err != nil {
		return iplib.Net6{}
	}
	if hostmask+netmask > 128 {
		fmt.Printf("total mask size cannot exceed 128 bits (%d + %d = %d)\n", netmask, hostmask, netmask+hostmask)
		os.Exit(1)
	}

	return iplib.NewNet6(address, netmask, hostmask)
}
