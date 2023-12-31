package cmd

import (
	"fmt"
	"os"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var nbContinue bool

var netBetweenCmd = &cobra.Command{
	Use:   "between <address> <address>",
	Short: "create a network between two IP addresses",
	Long: `
The 'net between' subcommand takes two IP addresses as arguments and returns
the largest IP netblock that will fit between them. Note that this might not
span the delta entirely. If the --continue flag is set then the command will
continue to return nets until the delta is spanned.

Examples:
  % ipfool net between 10.0.0.0 15.1.0.1
  10.0.0.0/7

  % ipfool net between --continue 10.0.0.0 15.1.0.1
  10.0.0.0/7
  12.0.0.0/7
  14.0.0.0/8
  15.0.0.0/16
  15.1.0.0/32
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(2),
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

		var iplast iplib.Net
		if iplib.EffectiveVersion(ipa) == iplib.IP4Version {
			iplast = iplib.Net4{}
		} else {
			iplast = iplib.Net6{}
		}

		for {
			ipnet, b, err := iplib.NewNetBetween(ipa, ipb)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(ipnet)

			if b == true || nbContinue == false {
				os.Exit(0)
			}

			if iplib.CompareIPs(ipnet.LastAddress(), ipb) >= 0 {
				os.Exit(0)
			}

			if iplast.IP() == nil {
				iplast = ipnet
			} else if iplib.CompareIPs(ipnet.IP(), iplast.IP()) > 0 {
				iplast = ipnet
			} else {
				os.Exit(0)
			}

			ipa = iplib.NextIP(ipnet.LastAddress())
			if iplib.CompareIPs(ipa, ipb) >= 0 {
				os.Exit(0)
			}
		}
	},
}

func init() {
	netRootCmd.AddCommand(netBetweenCmd)
	netBetweenCmd.Flags().BoolVarP(&nbContinue, "continue", "c", false, "keep going til no networks can be found")
}
