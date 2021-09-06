package cmd

import (
	"fmt"
	"github.com/c-robinson/iplib"
	"github.com/spf13/cobra"
)

var pnCIDR int

var netPrevCmd = &cobra.Command{
	Use:   "prevnet",
	Short: "get the previous netblock at the given mask (0 for same mask length)",
	Long: `
The prevnet subcommand takes a subnet as input and, by default, returns the
adjacent subnet preceding it, with the same mask length. The --cidr flag can
be used to generate the subnet at a different mask length but if the supplied
mask is larger than the original one it is very likely that the new subnet
will wind up being the supernet of the input subnet.`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"cidr"},
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)
		cidr, _ := ipnet.Mask().Size()
		if pnCIDR == 0 {
			pnCIDR = cidr
		}

		fmt.Printf("%-18s %-36s\n", "Original", ipnet.String())
		ViewIPAddress(ipnet)
		switch ipnet.Version() {
		case iplib.IP4Version:
			ipnets := ipnet.(iplib.Net4).PreviousNet(pnCIDR)
			fmt.Printf("%-18s %-36s\n", "Previous adjacent", ipnets.String())
			ViewIPAddress(ipnets)

		case iplib.IP6Version:
			ipnets := ipnet.(iplib.Net6).PreviousNet(pnCIDR)
			fmt.Printf("%-18s %-36s\n", "Previous adjacent", ipnets.String())
			ViewIPAddress(ipnets)
		}
	},
}

func init() {
	netRootCmd.AddCommand(netPrevCmd)
	netPrevCmd.Flags().IntVar(&pnCIDR, "cidr", 0, "new CIDR mask")
}
