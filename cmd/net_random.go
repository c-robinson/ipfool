package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var netRandomIPCmd = &cobra.Command{
	Use:   "random",
	Short: "return a random IP from the given subnet",
	Long: `
The 'net random' subcommand selects a random IP address from the given v4
or v6 subnet.`,
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
