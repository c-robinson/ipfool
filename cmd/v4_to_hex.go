package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v4ToHexCmd = &cobra.Command{
	Use:   "hex",
	Short: "dotted-decimal IPv4 address to hexadecimal",
	Long: `
The 'v4 to hex' command prints he hexadecimal representation of a given IPv4
address.`,
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
