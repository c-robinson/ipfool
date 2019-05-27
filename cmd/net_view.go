package cmd

import (
	"fmt"
	"github.com/c-robinson/iplib/iana"
	"github.com/spf13/cobra"
	"strings"

	"github.com/c-robinson/iplib"
)

var netViewCmd = &cobra.Command{
	Use:   "view",
	Short: "view details about an ipv4 or ipv6 netblock",
	Long:  "",
	DisableFlagsInUseLine: true,
	Args: cobra.ExactArgs(1),
	ValidArgs: []string{ "cidr" },
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)
		ViewIPAddress(ipnet)
	},
}

func ViewIPAddress(ipnet iplib.Net) {
	if ipnet.Version() == 4  {
		ViewIPv4Address(ipnet)
	} else {
		ViewIPv6Address(ipnet)
	}
}

func ViewIPv4Address(ipnet iplib.Net) {
	data := map[string]string{
		"Address": ipnet.IP.String(),
		"Netmask": iplib.HexStringToIP(ipnet.Mask.String()).String(),
		"Network": ipnet.NetworkAddress().String(),
		"First": ipnet.FirstAddress().String(),
		"Last": ipnet.LastAddress().String(),
		"Broadcast": ipnet.BroadcastAddress().String(),
		"Count": fmt.Sprintf("%d", ipnet.Count6()),
	}

	for _, k := range []string{"Address", "Netmask", "Network", "First", "Last", "Broadcast", "Count"} {
		fmt.Printf("%-18s %-16s\n", k, data[k])
	}
	rfclist := iana.GetRFCsForNetwork(ipnet)
	if len(rfclist) > 0 {
		fmt.Println("Registered in:", strings.Join(rfclist, ", "))
		fmt.Println(checkReservationBools(ipnet))
	}
}

func ViewIPv6Address(ipnet iplib.Net) {
	data := map[string]string{
		"Address": ipnet.IP.String(),
		"Netmask": putSeperatorsAroundIPv6Netmask(ipnet.Mask.String()),
		"First": iplib.ExpandIP6(ipnet.FirstAddress()),
		"Last": iplib.ExpandIP6(ipnet.LastAddress()),
		"Count": fmt.Sprintf("%d", ipnet.Count6()),
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
		if i%4 == 0 &&   i != 0 {
			ns += ":"
		}
		ns = ns + string(b)
	}
	return ns
}

func init() {
	netRootCmd.AddCommand(netViewCmd)
}
