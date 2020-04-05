package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var ipToBinaryCmd = &cobra.Command{
	Use:   "iptobinary",
	Short: "IPv4 or IPv6 address to binary",
	Long: `
The iptobinary subcommand prints a given IPv4 or IPv6 address as binary. This
can be useful for comparing netmask boundaries or as a calming meditative
exercise if you're kind of weird.`,
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
