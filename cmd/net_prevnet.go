package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var pnCIDR int

var netPrevCmd = &cobra.Command{
	Use:                   "prevnet",
	Short:                 "get the previous netblock at the given mask (0 for same mask length)",
	Long:                  "",
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"cidr"},
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)
		cidr, _ := ipnet.Mask.Size()
		if pnCIDR == 0 {
			pnCIDR = cidr
		}

		fmt.Printf("%-18s %-36s\n", "Original", ipnet.String())
		ViewIPAddress(ipnet)
		ipnets := ipnet.PreviousNet(pnCIDR)
		fmt.Printf("%-18s %-36s\n", "Previous adjacent", ipnets.String())

		ViewIPAddress(ipnets)

	},
}

func init() {
	netRootCmd.AddCommand(netPrevCmd)
	netPrevCmd.Flags().IntVar(&pnCIDR, "cidr", 0, "new CIDR mask")
}
