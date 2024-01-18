package cmd

import (
	"fmt"
	"strings"

	"github.com/c-robinson/iplib/v2/iana"
	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib/v2"
)

var netViewCmd = &cobra.Command{
	Use:   "view <network>",
	Short: "view details about an ipv4 or ipv6 netblock",
	Long: `
The 'net view' subcommand takes a subnet as input and prints some handy-dandy
info about it, like the network's first and last usable address, the number of
addresses it contains and whether all or part of the network overlaps with an
IANA reservation.

Note that the count that is returned is the number of *usable* addresses in
the block, not the total number. For IPv4 this means that in all but one case
the count will be 2 less than the total number of addresses in the block since
the first and last addresses are reserved for network and broadcast. The lone
exception is a /31, which has only two addresses and is only used in the wild
for numbering point-to-point links a la RFC 3021.

For information about the IANA registries see 'ipfool help iana'.

For information about the forwarding, private and reserved attributes, see
'ipfool help attributes'.

As for scope the rule is that 'view' reports "greedily" meaning that if any
part of the address block is named in an RFC (or has a boolean value of true
for any of the three special designations) then the entire block is reported
as such. So 'net view ::/0' will report the entire IPv6 space, every relevant
RFC and the most restrictive designation values.
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"cidr"},
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)
		viewIPAddress(ipnet)
	},
}

func viewIPAddress(ipnet iplib.Net) {
	switch ipnet.Version() {
	case iplib.IP4Version:
		netViewIPv4Address(ipnet.(iplib.Net4))

	case iplib.IP6Version:
		netViewIPv6Address(ipnet.(iplib.Net6))
	}
}

func netViewIPv4Address(ipnet iplib.Net4) {
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

func netViewIPv6Address(ipnet iplib.Net6) {
	size, _ := ipnet.Hostmask.Size()
	data := map[string]string{
		"Address":  ipnet.IP().String(),
		"Netmask":  putSeperatorsAroundIPv6Netmask(ipnet.Mask().String()),
		"First":    iplib.ExpandIP6(ipnet.FirstAddress()),
		"Last":     iplib.ExpandIP6(ipnet.LastAddress()),
		"Hostmask": putSeperatorsAroundIPv6Netmask(ipnet.Hostmask.String()),
		"Count":    fmt.Sprintf("%s", ipnet.Count().String()),
	}
	for _, k := range []string{"Address", "Netmask", "Hostmask", "First", "Last", "Count"} {
		if k == "Hostmask" && size == 0 {
			continue
		}
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
