package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
	"lukechampine.com/uint128"
)

var netEnumerateOffsetFlag, netEnumerateCountFlag int

var netEnumerateCmd = &cobra.Command{
	Use:   "enumerate <network>",
	Short: "print all IPs in the subnet (caveat emptor)",
	Long: `
'net enumerate' explicitly prints out all of the addresses in a given subnet,
one per line. This may take an astonishingly long time in the IPv6 case.

Flags:
 --count <int>    limit the number of IPs returned
 --offset <int>   start enumeration at this offset from the start address

Examples:
  % ipfool.go net enumerate --offset 15 --count 4 192.168.0.0/16
  192.168.0.15
  192.168.0.16
  192.168.0.17
  192.168.0.18
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)
		switch ipnet.Version() {
		case iplib.IP4Version:
			enumerateNet4(ipnet.(iplib.Net4))

		case iplib.IP6Version:
			enumerateNet6(ipnet.(iplib.Net6))
		}
	},
}

func enumerateNet4(n iplib.Net4) {
	var err error
	if netEnumerateOffsetFlag == 0 {
		netEnumerateOffsetFlag = 1
	}

	if netEnumerateCountFlag > int(n.Count())-netEnumerateOffsetFlag {
		fmt.Println("Requested count", netEnumerateCountFlag, "larger than block count", n.Count())
		return
	}

	if netEnumerateCountFlag == 0 {
		netEnumerateCountFlag = int(n.Count())
	}

	ip := iplib.IncrementIP4By(n.IP(), uint32(netEnumerateOffsetFlag))
	count := 0
	for {
		count++
		fmt.Println(ip)
		ip, err = n.NextIP(ip)
		if err != nil || count == netEnumerateCountFlag {
			return
		}
	}
}

func enumerateNet6(n iplib.Net6) {
	var err error
	z := uint128.From64(uint64(netEnumerateOffsetFlag))
	ip := iplib.IncrementIP6By(n.IP(), z)
	count := 0
	for {
		count++
		fmt.Println(ip)
		ip, err = n.NextIP(ip)
		if err != nil || count == netEnumerateCountFlag {
			return
		}
	}
}

func init() {
	netRootCmd.AddCommand(netEnumerateCmd)
	netEnumerateCmd.Flags().IntVarP(&netEnumerateOffsetFlag, "offset", "o", 0, "offset into the netblock to start from")
	netEnumerateCmd.Flags().IntVarP(&netEnumerateCountFlag, "count", "c", 0, "max number of entries to print")
}
