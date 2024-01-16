package cmd

import (
	"fmt"
	"math"
	"net"
	"os"
	"strconv"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
	"lukechampine.com/uint128"
)

var netDecrementByFlag string

var netDecrementCmd = &cobra.Command{
	Use:   "decrement <network> <address>",
	Short: "decrement an address within a netblock by an amount (default 1)",
	Long: `
'net decrement' decrements the provided IP so long as the result is still
within the provided netblock, or it exits with an error code 2. In the
IPv6 context the command will respect both netmask and hostmask boundaries if
provided.

Flags:
  --by <int>  decrement by this amount (default 1)

Examples:
  % ipfool net decrement 2001:db8::/56 2001:db8::fe --by 254
  2001:db8::

  % ipfool net decrement 2001:db8::/56:64 2001:db8:0:fe::  --by 254
  2001:db8::

  % ipfool net decrement 2001:db8::/56:64 2001:db8:0:fe::  --by 256
  cannot decrement by '256': result is outside of network
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(2),
	ValidArgs:             []string{"by"},
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)
		ip := retrieveIPAddress(args[1], v6)

		if !ipnet.Contains(ip) {
			fmt.Printf("'%s' is not contained within '%s'\n", ip, ipnet)
			os.Exit(2)
		}

		switch ipnet.Version() {
		case iplib.IP4Version:
			netDecrementIP4(ipnet.(iplib.Net4), ip)

		case iplib.IP6Version:
			netDecrementIP6(ipnet.(iplib.Net6), ip)
		}
	},
}

func netDecrementIP4(ipnet iplib.Net4, ip net.IP) {
	z, err := strconv.Atoi(netDecrementByFlag)
	if err != nil {
		fmt.Printf("cannot convert '%s' to an integer: %s\n", v6DecrementByFlag, err)
		os.Exit(1)
	}
	if z > math.MaxUint32 {
		fmt.Printf("cannot decrement by %d: too large\n", z)
		os.Exit(1)
	}
	xip := iplib.DecrementIP4By(ip, uint32(z))
	if ipnet.Contains(xip) {
		fmt.Println(xip)
	} else {
		fmt.Printf("cannot decrement by '%d': result is outside of network\n", z)
		os.Exit(2)
	}
}

func netDecrementIP6(ipnet iplib.Net6, ip net.IP) {
	z, err := uint128.FromString(netDecrementByFlag)
	if err != nil {
		fmt.Printf("cannot convert '%s' to an integer: %s\n", v6DecrementByFlag, err)
		os.Exit(1)
	}
	xip, err := iplib.DecrementIP6WithinHostmask(ip, ipnet.Hostmask, z)
	if err != nil {
		fmt.Printf("cannot decrement by '%s': result is outside of network\n", z)
		os.Exit(1)
	}
	if ipnet.Contains(xip) {
		fmt.Println(xip)
	} else {
		fmt.Printf("cannot decrement by '%s': result is outside of network\n", z)
		os.Exit(1)
	}
}

func init() {
	netRootCmd.AddCommand(netDecrementCmd)
	netDecrementCmd.Flags().StringVar(&netDecrementByFlag, "by", "1", "decrement address by count")
}
