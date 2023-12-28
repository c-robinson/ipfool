package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v6FromArpaCmd = &cobra.Command{
	Use:   "arpa",
	Short: "in-addr.arpa to IPv6 address",
	Long: `
The 'v6 from arpa' subcommand converts an address in in-addr.arpa format
to a regular-old IPv4 address. For example
1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa
becomes 2001:db8::1`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(iplib.ARPAToIP6(args[0]))
	},
}

func init() {
	v6FromCmd.AddCommand(v6FromArpaCmd)
}
