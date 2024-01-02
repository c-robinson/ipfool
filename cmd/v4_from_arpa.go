package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v4FromArpaCmd = &cobra.Command{
	Use:   "arpa <in-addr.arpa PTR record>",
	Short: "IPv4 address from DNS ARPA-domain PTR",
	Long: `
The 'v4 from arpa' subcommand converts an address in in-addr.arpa format
to a regular-old IPv4 address. 

Examples:
  % ipfool v4 from arpa 1.1.168.192.in-addr.arpa
  192.168.1.1
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(iplib.ARPAToIP4(args[0]))
	},
}

func init() {
	v4FromCmd.AddCommand(v4FromArpaCmd)
}
