package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var subCIDR int

var netSubnetCmd = &cobra.Command{
	Use:   "subnet",
	Short: "divide a netblock into subnets",
	Long:  "",
	DisableFlagsInUseLine: true,
	Args: cobra.ExactArgs(1),
	ValidArgs: []string{ "cidr" },
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)
		cidr, _ := ipnet.Mask.Size()
		if subCIDR == 0 {
			subCIDR = cidr + 1
		}
		if cidr > subCIDR {
			fmt.Printf("you appear to be trying to supernet (mask %d > supplied mask %d\n", cidr, subCIDR)
			os.Exit(1)
		}

		fmt.Printf("%-18s %-36s\n", "Original", ipnet.String())
		ViewIPAddress(ipnet)
		ipnets, _ := ipnet.Subnet(subCIDR)
		for i, ipn := range ipnets {
			fmt.Printf("%s %-11d %-36s\n", "Subnet", i+1, ipn.String())
			ViewIPAddress(ipn)
		}
	},
}

func init() {
	netRootCmd.AddCommand(netSubnetCmd)
	netSubnetCmd.Flags().IntVar(&subCIDR, "cidr", 0, "new CIDR mask")
}
