package cmd

import (
	"fmt"
	"github.com/c-robinson/iplib"
	"github.com/spf13/cobra"
)

var ip6ExpandCmd = &cobra.Command{
	Use:   "ip6expand",
	Short: "print full IPv6 address instead of the normal, concise format",
	Long:  "",
	DisableFlagsInUseLine: true,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v6)

		fmt.Println(iplib.ExpandIP6(ip))
	},
}

func init() {
	convertRootCmd.AddCommand(ip6ExpandCmd)
}
