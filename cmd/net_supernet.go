package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var superCIDR int

var netSupernetCmd = &cobra.Command{
	Use:                   "supernet",
	Short:                 "get the supernet of a given netblock",
	Long:                  "",
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"cidr"},
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)
		cidr, _ := ipnet.Mask.Size()
		if superCIDR == 0 {
			superCIDR = cidr - 1
		}
		if cidr < superCIDR {
			fmt.Printf("you appear to be trying to subnet (mask %d < supplied mask %d\n", cidr, superCIDR)
			os.Exit(1)
		}

		fmt.Printf("%-18s %-36s\n", "Original", ipnet.String())
		ViewIPAddress(ipnet)
		ipnets, _ := ipnet.Supernet(superCIDR)
		fmt.Printf("%-18s %-36s\n", "Supernet", ipnets.String())

		ViewIPAddress(ipnets)

	},
}

func init() {
	netRootCmd.AddCommand(netSupernetCmd)
	netSupernetCmd.Flags().IntVar(&superCIDR, "cidr", 0, "new CIDR mask")
}
