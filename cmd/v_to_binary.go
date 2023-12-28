package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib/v2"
)

var v46ToBinaryCmd = &cobra.Command{
	Use:   "binary",
	Short: "IPv4 or IPv6 address to binary",
	Long: `
The 'to binary' subcommand prints a given IPv4 or IPv6 address as binary. This
can be useful for comparing netmask boundaries or as a calming meditative
exercise if you're kind of weird.`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var xt addrType
		switch parent := cmd.Parent().Name(); parent {
		case "v4":
			xt = v4

		case "v6":
			xt = v6
		}

		ip := retrieveIPAddress(args[0], xt)

		fmt.Println(iplib.IPToBinaryString(ip))
	},
}

func init() {
	var v4ToBinaryCmd = *v46ToBinaryCmd
	v4ToCmd.AddCommand(&v4ToBinaryCmd)

	var v6ToBinaryCmd = *v46ToBinaryCmd
	v6ToCmd.AddCommand(&v6ToBinaryCmd)
}
