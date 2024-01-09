package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v4ToARPACmd = &cobra.Command{
	Use:   "arpa <address>",
	Short: "IPv4 address to DNS ARPA-domain PTR",
	Long: `
The 'v4 to arpa' command prints the in-addr.arpa entry for a given IPv4
address, suitable for use as a PTR record.

Examples:
  % ipfool v4 to arpa 192.168.1.1
  1.1.168.192.in-addr.arpa
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v4)
		fmt.Println(iplib.IP4ToARPA(ip))
	},
}

func init() {
	v4ToCmd.AddCommand(v4ToARPACmd)
}
