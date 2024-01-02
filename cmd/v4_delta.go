package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v4DeltaCmd = &cobra.Command{
	Use:   "delta <address1> <address2>",
	Short: "find the distance between two IPv4 addresses",
	Long: `
The 'v4 delta' subcommand takes two IPv4 addresses as input and returns an
integer of the difference between them.

Examples:
  % ipfool v4 delta 192.168.1.1 192.168.255.1
  65024
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ipa := retrieveIPAddress(args[0], v4)
		ipb := retrieveIPAddress(args[1], v4)
		fmt.Println(iplib.DeltaIP4(ipa, ipb))
	},
}

func init() {
	v4RootCmd.AddCommand(v4DeltaCmd)
}
