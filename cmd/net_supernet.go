package cmd

import (
	"fmt"
	"os"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var superCIDR int

var netSupernetCmd = &cobra.Command{
	Use:   "supernet",
	Short: "get the supernet of a given netblock",
	Long: `
The 'net supernet' subcommand takes a subnet as input and, by default, returns
that subnet's immediate parent (the subnet with a slightly greater netmask).
The --cidr flag can be used to request the supernet at a given mask, though it
is an error for the input mask to be smaller than the original (for which see
the subnet subcommand).`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"cidr"},
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)
		cidr, _ := ipnet.Mask().Size()
		if superCIDR == 0 {
			superCIDR = cidr - 1
		}
		if cidr < superCIDR {
			fmt.Printf("you appear to be trying to subnet (mask %d < supplied mask %d\n", cidr, superCIDR)
			os.Exit(1)
		}

		fmt.Printf("%-18s %-36s\n", "Original", ipnet.String())
		ViewIPAddress(ipnet)

		switch ipnet.Version() {
		case iplib.IP4Version:
			ipnets, _ := ipnet.(iplib.Net4).Supernet(superCIDR)
			fmt.Printf("%-18s %-36s\n", "Supernet", ipnets.String())
			ViewIPAddress(ipnets)

		case iplib.IP6Version:
			ipnets, _ := ipnet.(iplib.Net6).Supernet(superCIDR, 0)
			fmt.Printf("%-18s %-36s\n", "Supernet", ipnets.String())
			ViewIPAddress(ipnets)

		}
	},
}

func init() {
	netRootCmd.AddCommand(netSupernetCmd)
	netSupernetCmd.Flags().IntVar(&superCIDR, "cidr", 0, "new CIDR mask")
}
