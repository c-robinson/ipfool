package cmd

import (
	"fmt"
	"os"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var netNextCIDRFlag int
var netNextViewFlag bool

var netNextCmd = &cobra.Command{
	Use:   "nextnet <network>",
	Short: "get the next netblock at the given mask",
	Long: `
The 'net nextnet' subcommand takes a subnet as input and, by default, returns
the next subnet at the same make length. The --cidr flag can be used to request
the next net at a different mask.

Examples:
  % ipfool net nextnet 192.168.0.0/24
  192.168.1.0/24

  % ipfool net nextnet 2001:db8:a000::/35
  2001:db8:c000::/35

  % net nextnet --cidr 25 192.168.0.0/24
  192.168.1.0/25
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"cidr", "view"},
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)
		cidr, _ := ipnet.Mask().Size()
		if netNextCIDRFlag == 0 {
			netNextCIDRFlag = cidr
		}

		var ipnets iplib.Net
		switch ipnet.Version() {
		case iplib.IP4Version:
			ipnets = ipnet.(iplib.Net4).NextNet(netNextCIDRFlag)

		case iplib.IP6Version:
			ipnets = ipnet.(iplib.Net6).NextNet(netNextCIDRFlag)
		}

		if iplib.CompareNets(ipnets, ipnet) == 0 {
			fmt.Println("No next network")
			os.Exit(1)
		}

		if !netNextViewFlag {
			fmt.Println(ipnets.String())
			os.Exit(0)
		}

		fmt.Printf("%-18s %-36s\n", "Original", ipnet.String())
		ViewIPAddress(ipnet)

		fmt.Printf("%-18s %-36s\n", "Next adjacent", ipnets.String())
		ViewIPAddress(ipnets)
	},
}

func init() {
	netRootCmd.AddCommand(netNextCmd)
	netNextCmd.Flags().IntVar(&netNextCIDRFlag, "cidr", 0, "new CIDR mask")
	netNextCmd.Flags().BoolVar(&netNextViewFlag, "view", false, "get expanded view of subnets")
}
