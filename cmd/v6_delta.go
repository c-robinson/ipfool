package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v6DeltaCmd = &cobra.Command{
	Use:   "delta <address1> <address2>",
	Short: "find the distance between two IPv6 addresses",
	Long: `
The 'v6 delta' subcommand takes two IPv6 addresses as input and returns an
integer of the difference between them. 

Examples:
  % ipfool v6 delta 2001:db8:: 2001:db8::100:0     
  16777216
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ipa := retrieveIPAddress(args[0], v6)
		ipb := retrieveIPAddress(args[1], v6)
		fmt.Println(iplib.DeltaIP6(ipa, ipb))
	},
}

func init() {
	v6RootCmd.AddCommand(v6DeltaCmd)
}
