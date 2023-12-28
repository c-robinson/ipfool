package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v4FromArpaCmd = &cobra.Command{
	Use:   "arpa",
	Short: "in-addr.arpa to IPv4 address",
	Long: `
The 'v4 from arpa' subcommand converts an address in in-addr.arpa format
to a regular-old IPv4 address. For example 1.1.168.192.in-addr.arpa becomes
192.168.1.1`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(iplib.ARPAToIP4(args[0]))
	},
}

func init() {
	v4FromCmd.AddCommand(v4FromArpaCmd)
}
