package cmd

import (
	"fmt"
	"os"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var netSubnetCIDRFlag int
var netSubnetViewFlag bool

var netSubnetCmd = &cobra.Command{
	Use:   "subnet <network>",
	Short: "divide a netblock into subnets",
	Long: `
The 'net subnet' subcommand takes a subnet as input and, by default, splits
the subnet in half. It is an error to supply a mask larger than the input
mask (for which see the supernet subcommand).

Flags:
  --cidr <int>  next netblock at this mask length (default: same as input)
  --view        get expanded view of results

Examples:
  % ipfool net subnet 192.168.0.0/16
  192.168.0.0/17
  192.168.128.0/17

  % ipfool net subnet 2001:db8::/32
  2001:db8::/33
  2001:db8:8000::/33

  % ipfool net subnet --cidr 35 2001:db8::/32
  2001:db8::/35
  2001:db8:2000::/35
  2001:db8:4000::/35
  2001:db8:6000::/35
  2001:db8:8000::/35
  2001:db8:a000::/35
  2001:db8:c000::/35
  2001:db8:e000::/35
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"cidr", "view"},
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)
		cidr, _ := ipnet.Mask().Size()
		if netSubnetCIDRFlag == 0 {
			netSubnetCIDRFlag = cidr + 1
		}
		if cidr > netSubnetCIDRFlag {
			fmt.Printf("you appear to be trying to supernet (mask %d > supplied mask %d\n", cidr, netSubnetCIDRFlag)
			os.Exit(1)
		}

		switch ipnet.Version() {
		case iplib.IP4Version:
			ipnets, _ := ipnet.(iplib.Net4).Subnet(netSubnetCIDRFlag)
			if !netSubnetViewFlag {
				for _, ipn := range ipnets {
					fmt.Println(ipn.String())
				}
				os.Exit(0)
			}

			fmt.Printf("%-18s %-36s\n", "Original", ipnet.String())
			ViewIPAddress(ipnet)
			for i, ipn := range ipnets {
				fmt.Printf("%s %-11d %-36s\n", "Subnet", i+1, ipn.String())
				ViewIPAddress(ipn)
			}

		case iplib.IP6Version:
			ipnets, _ := ipnet.(iplib.Net6).Subnet(netSubnetCIDRFlag, 0)
			if !netSubnetViewFlag {
				for _, ipn := range ipnets {
					fmt.Println(ipn.String())
				}
				return
			}

			ViewIPAddress(ipnet)
			for i, ipn := range ipnets {
				fmt.Printf("%s %-11d %-36s\n", "Subnet", i+1, ipn.String())
				ViewIPAddress(ipn)
			}

		}
	},
}

func init() {
	netRootCmd.AddCommand(netSubnetCmd)
	netSubnetCmd.Flags().IntVar(&netSubnetCIDRFlag, "cidr", 0, "new CIDR mask")
	netSubnetCmd.Flags().BoolVar(&netSubnetViewFlag, "view", false, "get expanded view of subnets")
}
