package cmd

import (
	"fmt"
	"strings"

	"github.com/c-robinson/iplib/iana/v2"
	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib/v2"
)

var netViewCmd = &cobra.Command{
	Use:   "view",
	Short: "view details about an ipv4 or ipv6 netblock",
	Long: `
The 'net view' subcommand takes a subnet as input and prints some handy-dandy
info about it, like the network's first and last usable address, the number of
addresses it contains, whether all or part of the network overlap with an
IANA reservation (such as the RFC 1918 private IPv4 networks or the RFC 3849
IPv6 address space set aside for use in documentation.

Note that the count that is returned is the number of *usable* addresses in
the block, not the total number. For IPv4 this means that in all but one case
the count will be 2 less than the total number of addresses in the block since
the first and last addresses are reserved for network and broadcast. The lone
exception is a /31, which has only two addresses and is only used in the wild
for numbering point-to-point links a la RFC 3021.`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"cidr"},
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)
		ViewIPAddress(ipnet)
	},
}

func ViewIPAddress(ipnet iplib.Net) {
	switch ipnet.Version() {
	case iplib.IP4Version:
		ViewIPv4Address(ipnet.(iplib.Net4))

	case iplib.IP6Version:
		ViewIPv6Address(ipnet.(iplib.Net6))
	}
}

func ViewIPv4Address(ipnet iplib.Net4) {
	data := map[string]string{
		"Address":   ipnet.IP().String(),
		"Netmask":   iplib.HexStringToIP(ipnet.Mask().String()).String(),
		"Network":   ipnet.NetworkAddress().String(),
		"First":     ipnet.FirstAddress().String(),
		"Last":      ipnet.LastAddress().String(),
		"Wildcard":  ipnet.Wildcard().String(),
		"Broadcast": ipnet.BroadcastAddress().String(),
		"Count":     fmt.Sprintf("%d", ipnet.Count()),
	}

	for _, k := range []string{"Address", "Netmask", "Network", "First", "Last", "Wildcard", "Broadcast", "Count"} {
		fmt.Printf("%-18s %-16s\n", k, data[k])
	}
	rfclist := iana.GetRFCsForNetwork(ipnet)
	if len(rfclist) > 0 {
		fmt.Println("Registered in:", strings.Join(rfclist, ", "))
		fmt.Println(checkReservationBools(ipnet))
	}
}

func ViewIPv6Address(ipnet iplib.Net6) {
	data := map[string]string{
		"Address": ipnet.IP().String(),
		"Netmask": putSeperatorsAroundIPv6Netmask(ipnet.Mask().String()),
		"First":   iplib.ExpandIP6(ipnet.FirstAddress()),
		"Last":    iplib.ExpandIP6(ipnet.LastAddress()),
		"Count":   fmt.Sprintf("%s", ipnet.Count().String()),
	}

	for _, k := range []string{"Address", "Netmask", "First", "Last", "Count"} {
		fmt.Printf("%-18s %-16s\n", k, data[k])
	}
	rfclist := iana.GetRFCsForNetwork(ipnet)
	if len(rfclist) > 0 {
		fmt.Println("Registered in:", strings.Join(rfclist, ", "))
		fmt.Println(checkReservationBools(ipnet))
	}
}

func checkReservationBools(ipnet iplib.Net) string {
	s := "Network "
	if iana.IsForwardable(ipnet) {
		s = s + "may be forwarded, "
	} else {
		s = s + "may not be forwarded, "
	}

	if iana.IsGlobal(ipnet) {
		s = s + "is globally accessible, "
	} else {
		s = s + "is private, "
	}

	if iana.IsReserved(ipnet) {
		s = s + "is reserved\n"
	} else {
		s = s + "is not reserved\n"
	}
	return s
}

func putSeperatorsAroundIPv6Netmask(s string) string {
	var ns string
	for i, b := range s {
		if i%4 == 0 && i != 0 {
			ns += ":"
		}
		ns = ns + string(b)
	}
	return ns
}

func init() {
	netRootCmd.AddCommand(netViewCmd)
}
