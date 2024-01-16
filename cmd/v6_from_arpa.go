package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v6FromArpaCmd = &cobra.Command{
	Use:   "arpa <ip6.arpa PTR record>",
	Short: "IPv6 address from DNS ARPA-domain PTR",
	Long: `
'v6 from arpa' converts an address in ip6.arpa format to a regular-old IPv6
address.

Examples:
  % ipfool v6 from arpa 1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa
  2001:db8::1
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(iplib.ARPAToIP6(args[0]))
	},
}

func init() {
	v6FromCmd.AddCommand(v6FromArpaCmd)
}
