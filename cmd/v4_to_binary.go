package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib/v2"
)

var v4ToBinaryCmd = &cobra.Command{
	Use:   "binary <address>",
	Short: "IPv4 address to binary",
	Long: `
The 'v4 to binary' subcommand prints a given IPv4 address as binary. This can
be useful for comparing netmask boundaries or as a substitute for counting
sheep if you're sleepy.

Examples:
  % ipfool v4 to binary 192.168.1.1
  11000000.10101000.00000001.00000001
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v4)
		fmt.Println(iplib.IPToBinaryString(ip))
	},
}

func init() {
	v4ToCmd.AddCommand(v4ToBinaryCmd)
}
