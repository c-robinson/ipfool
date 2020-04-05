package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var ipToARPACmd = &cobra.Command{
	Use:   "iptoarpa",
	Short: "IPv4 or IPv6 address to DNS ARPA-domain PTR",
	Long: `
The iptoarpa command prints the in-addr.arpa entry for a given IPv4 or IPv6
address, suitible for use as a PTR record.`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v46)

		fmt.Println(iplib.IPToARPA(ip))
	},
}

func init() {
	convertRootCmd.AddCommand(ipToARPACmd)
}
