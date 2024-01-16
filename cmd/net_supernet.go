package cmd

import (
	"fmt"
	"os"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var netSupernetCIDRFlag int
var netSupernetViewFlag bool

var netSupernetCmd = &cobra.Command{
	Use:   "supernet <network>",
	Short: "get the supernet of a given netblock",
	Long: `
'net supernet' takes a subnet as input and, by default, returns that subnet's
immediate parent (the subnet with a slightly greater netmask). It is an error
for the input mask to be smaller than the original (for which see 
'net subnet').

Flags:
  --cidr <int>  next netblock at this mask length (default: same as input)
  --view        get expanded view of results

Examples:
  % ipfool net supernet 192.168.1.0/17
  192.168.0.0/16

  % ipfool net supernet 2001:db8:a000::/35
  2001:db8:8000::/34

  % ipfool net supernet --cidr 32 2001:db8:a000::/35
  2001:db8::/32
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"cidr", "view"},
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)
		cidr, _ := ipnet.Mask().Size()
		if netSupernetCIDRFlag == 0 {
			netSupernetCIDRFlag = cidr - 1
		}
		if cidr < netSupernetCIDRFlag {
			fmt.Printf("you appear to be trying to subnet (mask %d < supplied mask %d\n", cidr, netSupernetCIDRFlag)
			os.Exit(1)
		}

		if cidr == 0 {
			fmt.Println("Cannot supernet a /0")
			os.Exit(1)
		}

		var ipnets iplib.Net
		switch ipnet.Version() {
		case iplib.IP4Version:
			ipnets, _ = ipnet.(iplib.Net4).Supernet(netSupernetCIDRFlag)

		case iplib.IP6Version:
			ipnets, _ = ipnet.(iplib.Net6).Supernet(netSupernetCIDRFlag, 0)

		}
		if !netSupernetViewFlag {
			fmt.Println(ipnets.String())
			os.Exit(0)
		}
		fmt.Printf("%-18s %-36s\n", "Original", ipnet.String())
		ViewIPAddress(ipnet)

		fmt.Printf("%-18s %-36s\n", "Supernet", ipnets.String())
		ViewIPAddress(ipnets)

	},
}

func init() {
	netRootCmd.AddCommand(netSupernetCmd)
	netSupernetCmd.Flags().IntVar(&netSupernetCIDRFlag, "cidr", 0, "new CIDR mask")
	netSupernetCmd.Flags().BoolVar(&netSupernetViewFlag, "view", false, "get expanded view of subnets")

}
