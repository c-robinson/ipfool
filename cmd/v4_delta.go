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
'v4 delta' returns the numerical difference between two supplied addresses.

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
