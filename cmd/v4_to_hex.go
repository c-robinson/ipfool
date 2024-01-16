package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v4ToHexCmd = &cobra.Command{
	Use:   "hex <address>",
	Short: "dotted-decimal IPv4 address to hexadecimal",
	Long: `
'v4 to hex' prints he hexadecimal representation of a given IPv4 address.

Examples:
  % ipfool v4 to hex 192.168.1.1
  c0a80101
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v4)
		fmt.Println(iplib.IPToHexString(ip))
	},
}

func init() {
	v4ToCmd.AddCommand(v4ToHexCmd)
}
