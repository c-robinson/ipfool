package cmd

import (
	"fmt"
	"github.com/c-robinson/iplib"
	"github.com/spf13/cobra"
)

var ipToIntCmd = &cobra.Command{
	Use:   "iptoint",
	Short: "IPv4 or IPv6 address to integer",
	Long: `
The iptoint subcommand converts a given IPv4 or IPv6 command into an integer
where the first address (0.0.0.0 or ::) are indexed to 0.`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v46)

		if iplib.Version(ip) == 4 {
			fmt.Println(iplib.IP4ToUint32(ip))
		} else {
			fmt.Println(iplib.IPToBigint(ip))
		}
	},
}

func init() {
	convertRootCmd.AddCommand(ipToIntCmd)
}
