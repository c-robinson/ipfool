package cmd

import (
	"bufio"
	"fmt"
	"io"
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

// genericNetIncrement is a helper function to increment an IP address within
// a network regardless of IP version
func genericNetIncrement(ip net.IP, ipnet iplib.Net) (net.IP, error) {
	switch ipnet.Version() {
	case 4:
		return ipnet.(iplib.Net4).NextIP(ip)

	case 6:
		return ipnet.(iplib.Net6).NextIP(ip)
	}
	return nil, fmt.Errorf("unknown IP version")
}

// ipListPruner takes a []net.IP and an iplib.Net and removes all entries from
// the ip array that are not contained within the Net
func ipListPruner(n iplib.Net, iplist []net.IP) []net.IP {
	var newlist []net.IP
	for _, ip := range iplist {
		if n.Contains(ip) {
			newlist = append(newlist, ip)
		}
	}
	return newlist
}

// ipListSplitter splits input strings on commas, spaces and newlines. It is
// used as a custom SplitFunc by bufio.Scanner in retrieveIPList.
func ipListSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == ',' || data[i] == ' ' || data[i] == '\n' {
			return i + 1, data[0:i], nil
		}
	}
	if !atEOF {
		return 0, nil, nil
	}
	return len(data), data, bufio.ErrFinalToken
}

// ipListUniquer takes a []net.IP and deduplicates entries in it
func ipListUniquer(iplist []net.IP) []net.IP {
	seen := make(map[string]bool)
	uniq := []net.IP{}
	for _, ip := range iplist {
		if _, ok := seen[ip.String()]; !ok {
			seen[ip.String()] = true
			uniq = append(uniq, ip)
		}
	}
	return uniq
}

// respondToTrueFalseQuestion takes a result and a bool indicating whether to
// print the result and exits with the appropriate exit code
func respondToTrueFalseQuestion(result, exitCodeOnly bool) {
	if !exitCodeOnly {
		fmt.Printf("%t\n", result)
	}
	if result {
		os.Exit(0)
	}
	os.Exit(2)
}

// retrieveIPAddress takes a string and an addrType and returns a net.IP. It
// first attempts to parse the string as an IP address. If that fails it exits
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

// retrieveIPList takes the cobra command args array and the io.Reader provided
// by cmd.InOrStdin() and returns a list of net.IP addresses. It first checks
// to see if --file was specified and if so opens the file. If not and if args
// contains the single element "-" then it uses the io.Reader. If neither of
// those are true, it returns an error.
//
// if stdin is the method of input retrieveIPList sets the boolean to true
// in the return value
func retrieveIPList(args []string, stdin io.Reader) ([]net.IP, bool, error) {
	var inputReader io.Reader
	iplist := []net.IP{}
	isStdin := false

	if listFileFlag != "" {
		file, err := os.Open(listFileFlag)
		if err != nil {
			return iplist, isStdin, fmt.Errorf("failed open file: %v", err)
		}
		inputReader = file
	} else if len(args) > 0 && args[len(args)-1] == "-" {
		inputReader = stdin
		isStdin = true
	} else {
		return iplist, isStdin, fmt.Errorf("no input file (--file <filename>) or stdin (-) specified")
	}

	scanner := bufio.NewScanner(inputReader)
	scanner.Split(ipListSplitter)

	for scanner.Scan() {
		ip := net.ParseIP(scanner.Text())
		if ip != nil {
			iplist = append(iplist, ip)
		}
	}
	return iplist, isStdin, nil
}

// retrieveIPNetwork takes a string and an addrType and returns an iplib.Net.
// It first attempts to parse the string as an IP network. If that fails it
// exits
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

// retrieveIPNetwork6WithHostmask takes a string and returns an iplib.Net6. It
// first attempts to parse the string as an IP network with hostmask. If that
// fails it returns an empty iplib.Net6
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
