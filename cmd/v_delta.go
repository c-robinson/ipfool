package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v46DeltaCmd = &cobra.Command{
	Use:   "delta",
	Short: "find the distance between two IP addresses",
	Long: `
The 'delta' subcommand takes two IP addresses as input and returns an
integer of the difference between them. For example the delta between
192.168.0.0 and 192.168.0.255 is 255`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		switch parent := cmd.Parent().Name(); parent {
		case "v4":
			ipa := retrieveIPAddress(args[0], v4)
			ipb := retrieveIPAddress(args[1], v4)
			fmt.Println(iplib.DeltaIP4(ipa, ipb))

		case "v6":
			ipa := retrieveIPAddress(args[0], v6)
			ipb := retrieveIPAddress(args[1], v6)
			fmt.Println(iplib.DeltaIP6(ipa, ipb))

		}
	},
}

func init() {
	var v4DeltaCmd = *v46DeltaCmd
	var v6DeltaCmd = *v46DeltaCmd
	v4RootCmd.AddCommand(&v4DeltaCmd)
	v6RootCmd.AddCommand(&v6DeltaCmd)
}
