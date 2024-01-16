package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v6ToARPACmd = &cobra.Command{
	Use:   "arpa <address>",
	Short: "IPv6 address to DNS ARPA-domain PTR",
	Long: `
'v6 to arpa' converts an IPv6 address to an ip6.arpa address, suitable for use
as a PTR record.

Examples:
  % ipfool v6 to arpa 2001:db8::1
  1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v6)
		fmt.Println(iplib.IP6ToARPA(ip))
	},
}

func init() {
	v6ToCmd.AddCommand(v6ToARPACmd)
}
