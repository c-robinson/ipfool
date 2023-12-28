package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v6ExpandCmd = &cobra.Command{
	Use:   "expand",
	Short: "print full IPv6 address instead of the normal, concise format",
	Long: `
The 'v6 expand' subcommand prints an IPv6 address without the shortcuts that
are often used to condense the address by removing obvious-from-context
zeroes. So 2001:db8:: becomes 2001:0db8:0000:0000:0000:0000:0000:0000.`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v6)

		fmt.Println(iplib.ExpandIP6(ip))
	},
}

func init() {
	v6RootCmd.AddCommand(v6ExpandCmd)
}
