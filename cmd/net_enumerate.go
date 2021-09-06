package cmd

import (
	"fmt"
	"github.com/c-robinson/iplib"
	"github.com/spf13/cobra"
	"math/big"
)

var neOffset, neCount int


var netEnumerateCmd = &cobra.Command{
	Use:   "enumerate",
	Short: "print all IPs in the subnet (caveat emptor)",
	Long: `
The enumerate subcommand explicitly prints out all of the addresses in a
given subnet, one per line. This may take an astonishingly long time in the
IPv6 case.`,
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
	if neOffset == 0 {
		neOffset = 1
	}

	if neCount > int(n.Count()) - neOffset {
		fmt.Println("Requested count", neCount, "larger than block count", n.Count())
		return
	}

	if neCount == 0 {
		neCount = int(n.Count())
	}

	ip := iplib.IncrementIP4By(n.IP(), uint32(neOffset))
	count := 0
	for {
		count++
		fmt.Println(ip)
		ip, err = n.NextIP(ip)
		if err != nil || count == neCount {
			return
		}
	}
}

func enumerateNet6(n iplib.Net6) {
	var err error
	z := big.NewInt(int64(neOffset))
	ip := iplib.IncrementIP6By(n.IP(), z)
	count := 0
	for {
		count++
		fmt.Println(ip)
		ip, err = n.NextIP(ip)
		if err != nil || count == neCount {
			return
		}
	}
}

func init() {
	netRootCmd.AddCommand(netEnumerateCmd)
	netEnumerateCmd.Flags().IntVarP(&neOffset, "offset", "o",0, "offset into the netblock to start from")
	netEnumerateCmd.Flags().IntVarP(&neCount, "count", "c",0, "max number of entries to print")
}
