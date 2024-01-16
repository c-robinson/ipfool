package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v4ToIntCmd = &cobra.Command{
	Use:   "int <address>",
	Short: "IPv4 address to 32bit unsigned integer",
	Long: `
'v4 to int' converts a given IPv4 address into an integer where the first 
address (0.0.0.0) is indexed to 0.

Examples:
  % ipfool v4 to int 192.168.1.1
  3232235777
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v4)
		fmt.Println(iplib.IP4ToUint32(ip))
	},
}

func init() {
	v4ToCmd.AddCommand(v4ToIntCmd)
}
