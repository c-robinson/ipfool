package cmd

import (
	"fmt"
	"os"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
	"lukechampine.com/uint128"
)

var v6DecrementByFlag string

var v6DecrementCmd = &cobra.Command{
	Use:   "decrement <address>",
	Short: "decrement an IPv6 address by a given amount (default 1)",
	Long: `
The 'v6 decrement' command takes an IPv6 address as input. If no arguments
are given it will decrement the address by one.

Flags:
  --by <int>  decrement by this amount (default 1)

Examples:
  % ipfool v6 decrement 2001:db8::1
  2001:db8::

  % ipfool v6 decrement --by 16777216 2001:db8::100:0
  2001:db8::
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"by"},
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v6)
		z, err := uint128.FromString(v6DecrementByFlag)
		if err != nil {
			fmt.Printf("cannot convert '%s' to an integer: %s\n", v6DecrementByFlag, err)
			os.Exit(1)
		}
		fmt.Println(iplib.DecrementIP6By(ip, z))
	},
}

func init() {
	v6RootCmd.AddCommand(v6DecrementCmd)
	v6DecrementCmd.Flags().StringVar(&v6DecrementByFlag, "by", "1", "decrement address by count")

}
