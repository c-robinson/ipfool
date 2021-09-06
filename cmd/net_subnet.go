package cmd

import (
	"fmt"
	"github.com/c-robinson/iplib"
	"github.com/spf13/cobra"
	"os"
)

var subCIDR int

var netSubnetCmd = &cobra.Command{
	Use:   "subnet",
	Short: "divide a netblock into subnets",
	Long: `
The subnet subcommand takes a subnet as input and, by default, splits the
subnet in half. The --cidr command can be supplied to divide it differently.
It is an error to supply a mask larger than the input mask (for which see the
supernet subcommand).`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"cidr"},
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)
		cidr, _ := ipnet.Mask().Size()
		if subCIDR == 0 {
			subCIDR = cidr + 1
		}
		if cidr > subCIDR {
			fmt.Printf("you appear to be trying to supernet (mask %d > supplied mask %d\n", cidr, subCIDR)
			os.Exit(1)
		}

		switch ipnet.Version() {
		case iplib.IP4Version:
			fmt.Printf("%-18s %-36s\n", "Original", ipnet.String())
			ViewIPAddress(ipnet)
			ipnets, _ := ipnet.(iplib.Net4).Subnet(subCIDR)
			for i, ipn := range ipnets {
				fmt.Printf("%s %-11d %-36s\n", "Subnet", i+1, ipn.String())
				ViewIPAddress(ipn)
			}

		case iplib.IP6Version:
			ViewIPAddress(ipnet)
			ipnets, _ := ipnet.(iplib.Net6).Subnet(subCIDR, 0)
			for i, ipn := range ipnets {
				fmt.Printf("%s %-11d %-36s\n", "Subnet", i+1, ipn.String())
				ViewIPAddress(ipn)
			}

		}
	},
}

func init() {
	netRootCmd.AddCommand(netSubnetCmd)
	netSubnetCmd.Flags().IntVar(&subCIDR, "cidr", 0, "new CIDR mask")
}
