package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib/v2"
)

var v6ToBinaryCmd = &cobra.Command{
	Use:   "binary <address>",
	Short: "IPv6 address to binary",
	Long: `
'v6 to binary' prints a given IPv6 address as binary. This can be useful for
comparing netmask boundaries or as a calming meditative exercise if you're
kind of weird.

Examples:
  % ipfool v6 to binary 2001:db8::1
  00100000.00000001.00001101.10111000.00000000.00000000.00000000.00000000.00000000.00000000.00000000.00000000.00000000.00000000.00000000.00000001
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], 6)
		fmt.Println(iplib.IPToBinaryString(ip))
	},
}

func init() {
	v6ToCmd.AddCommand(v6ToBinaryCmd)
}
