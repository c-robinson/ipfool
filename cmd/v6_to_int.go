package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v6ToIntCmd = &cobra.Command{
	Use:   "int",
	Short: "IPv6 address to 128bit integer",
	Long: `
The 'v6 to int' subcommand converts a given IPv6 address into an integer
where the first address (::) is indexed to 0.`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v6)
		fmt.Println(iplib.IPToBigint(ip).String())
	},
}

func init() {
	v6ToCmd.AddCommand(v6ToIntCmd)
}
