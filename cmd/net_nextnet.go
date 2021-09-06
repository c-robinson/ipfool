package cmd

import (
	"fmt"
	"github.com/c-robinson/iplib"
	"github.com/spf13/cobra"
)

var nnCIDR int

var netNextCmd = &cobra.Command{
	Use:   "nextnet",
	Short: "get the next netblock at the given mask (0 for same mask length)",
	Long: `
The nextnet subcommand takes a subnet as input and, by default, returns the
next subnet at the same make length. The --cidr flag can be used to request
the next net at a different mask.`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"cidr"},
	Run: func(cmd *cobra.Command, args []string) {
		_, ipnet, _ := iplib.ParseCIDR(args[0])
		cidr, _ := ipnet.Mask().Size()
		if nnCIDR == 0 {
			nnCIDR = cidr
		}

		fmt.Printf("%-18s %-36s\n", "Original", ipnet.String())
		ViewIPAddress(ipnet)

		switch ipnet.Version() {
		case iplib.IP4Version:
			ipnets := ipnet.(iplib.Net4).NextNet(pnCIDR)
			fmt.Printf("%-18s %-36s\n", "Next adjacent", ipnets.String())
			ViewIPAddress(ipnets)

		case iplib.IP6Version:
			ipnets := ipnet.(iplib.Net6).NextNet(pnCIDR)
			fmt.Printf("%-18s %-36s\n", "Next adjacent", ipnets.String())
			ViewIPAddress(ipnets)

		}
	},
}

func init() {
	netRootCmd.AddCommand(netNextCmd)
	netNextCmd.Flags().IntVar(&nnCIDR, "cidr", 0, "new CIDR mask")
}
