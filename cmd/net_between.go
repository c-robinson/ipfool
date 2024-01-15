package cmd

import (
	"fmt"
	"os"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var netBetweenViewFlag bool

var netBetweenCmd = &cobra.Command{
	Use:   "between <address> <address>",
	Short: "create a network between two IP addresses",
	Long: `
The 'net between' subcommand takes two IP addresses as arguments and returns
a list of the netblocks required to span them, inclusive of the first address
and exclusive of the last.

Flags:
  --view  get expanded view of results

Examples:
  % ipfool net between 10.0.0.0 15.1.0.1
  10.0.0.0/7
  12.0.0.0/7
  14.0.0.0/8
  15.0.0.0/16
  15.1.0.0/32
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(2),
	ValidArgs:             []string{"continue"},
	Run: func(cmd *cobra.Command, args []string) {
		ipa := retrieveIPAddress(args[0], v46)
		ipb := retrieveIPAddress(args[1], v46)

		if iplib.EffectiveVersion(ipa) != iplib.EffectiveVersion(ipb) {
			fmt.Println("mismatched IP versions")
			os.Exit(1)
		}

		if iplib.CompareIPs(ipa, ipb) > 0 {
			ipa, ipb = ipb, ipa
		}

		ipnets, err := iplib.AllNetsBetween(ipa, ipb)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, ipnet := range ipnets {
			fmt.Println(ipnet)
		}
	},
}

func init() {
	netRootCmd.AddCommand(netBetweenCmd)
	netBetweenCmd.Flags().BoolVar(&netBetweenViewFlag, "view", false, "get expanded view of subnets")

}
