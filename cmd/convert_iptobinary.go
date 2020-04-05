package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var ipToBinaryCmd = &cobra.Command{
	Use:                   "iptobinary",
	Short:                 "IPv4 or IPv6 address to binary",
	Long:                  "",
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v46)

		fmt.Println(iplib.IPToBinaryString(ip))
	},
}

func init() {
	convertRootCmd.AddCommand(ipToBinaryCmd)
}
