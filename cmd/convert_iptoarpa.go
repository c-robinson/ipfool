package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var ipToARPACmd = &cobra.Command{
	Use:                   "iptoarpa",
	Short:                 "IPv4 or IPv6 address to DNS ARPA-domain PTR",
	Long:                  "",
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
