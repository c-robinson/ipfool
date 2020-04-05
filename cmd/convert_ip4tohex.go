package cmd

import (
	"fmt"
	"github.com/c-robinson/iplib"
	"github.com/spf13/cobra"
)

var ip4ToHexCmd = &cobra.Command{
	Use:                   "ip4tohex",
	Short:                 "dotted-decimal IPv4 address to hexadecimal",
	Long:                  "",
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v4)

		fmt.Println(iplib.IPToHexString(ip))
	},
}

func init() {
	convertRootCmd.AddCommand(ip4ToHexCmd)
}
