package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v46ToARPACmd = &cobra.Command{
	Use:   "arpa",
	Short: "IP address to DNS ARPA-domain PTR",
	Long: `
The 'to arpa' command prints the in-addr.arpa entry for a given IPv4 or IPv6
address, suitable for use as a PTR record.`,
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
		fmt.Println(iplib.IPToARPA(ip))
	},
}

func init() {
	var v4ToARPACmd = *v46ToARPACmd
	v4ToCmd.AddCommand(&v4ToARPACmd)

	var v6ToARPACmd = *v46ToARPACmd
	v6ToCmd.AddCommand(&v6ToARPACmd)
}
