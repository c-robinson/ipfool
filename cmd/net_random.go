package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var netRandomIPCmd = &cobra.Command{
	Use:   "random <network>",
	Short: "return a random IP from the given netblock",
	Long: `
'net random' selects a random IP address from the given v4 or v6 subnet.

Examples:
  % ipfool net random 192.168.0.0/16
  192.168.4.164

  % ipfool net random 2001:db8::/64
  2001:db8::8a48:1da:8b4b:d7d1
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)

		switch ipnet.Version() {
		case iplib.IP4Version:

			fmt.Println(ipnet.(iplib.Net4).RandomIP())

		case iplib.IP6Version:
			fmt.Println(ipnet.(iplib.Net6).RandomIP())
		}
	},
}

func init() {
	netRootCmd.AddCommand(netRandomIPCmd)
}
