package cmd

import (
	"fmt"
	"os"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var netPrevCIDRFlag int
var netPrevViewFlag bool

var netPrevCmd = &cobra.Command{
	Use:   "prevnet <network>",
	Short: "get the previous netblock at the given mask",
	Long: `
The 'net prevnet' subcommand takes a subnet as input and, by default, returns
the adjacent subnet preceding it, with the same mask length. If the --cidr flag
is supplied with a larger mask than the original one it is very likely that
the new subnet will wind up being the supernet of the input subnet.

Flags:
  --cidr <int>  next netblock at this mask length (default: same as input)
  --view        get expanded view of results

Examples:
  % ipfool net prevnet 192.168.1.0/24
  192.168.0.0/24

  % ipfool net prevnet 2001:db8:e000::/35
  2001:db8:c000::/35

  % ipfool net prevnet --cidr 64 2001:db8:e000::/35
  2001:db8:dfff:ffff::/64
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"cidr", "view"},
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)
		cidr, _ := ipnet.Mask().Size()
		if netPrevCIDRFlag == 0 {
			netPrevCIDRFlag = cidr
		}

		var ipnets iplib.Net
		switch ipnet.Version() {
		case iplib.IP4Version:
			ipnets = ipnet.(iplib.Net4).PreviousNet(netPrevCIDRFlag)

		case iplib.IP6Version:
			ipnets = ipnet.(iplib.Net6).PreviousNet(netPrevCIDRFlag)
		}

		if iplib.CompareNets(ipnets, ipnet) == 0 {
			fmt.Println("No previous network")
			os.Exit(1)
		}

		if !netPrevViewFlag {
			fmt.Println(ipnets.String())
			os.Exit(0)
		}

		fmt.Printf("%-18s %-36s\n", "Original", ipnet.String())
		ViewIPAddress(ipnet)

		fmt.Printf("%-18s %-36s\n", "Previous adjacent", ipnets.String())
		ViewIPAddress(ipnets)
	},
}

func init() {
	netRootCmd.AddCommand(netPrevCmd)
	netPrevCmd.Flags().IntVar(&netPrevCIDRFlag, "cidr", 0, "new CIDR mask")
	netPrevCmd.Flags().BoolVar(&netPrevViewFlag, "view", false, "get expanded view of subnets")
}
