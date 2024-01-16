package cmd

import (
	"fmt"
	"os"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
	"lukechampine.com/uint128"
)

var v6IncrementByFlag string

var v6IncrementCmd = &cobra.Command{
	Use:   "increment <address>",
	Short: "increment an IPv6 address by a given amount (default 1)",
	Long: `
'v6 increment' increments the provided address by the given amount, or 1 if
no amount is provided.

Flags:
  --by <int>  increment by this amount (default 1)

Examples:
  % ipfool v6 increment 2001:db8::
  2001:db8::1

  % ipfool v6 increment --by 16777216 2001:db8::
  2001:db8::100:0
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"by"},
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v6)
		z, err := uint128.FromString(v6IncrementByFlag)
		if err != nil {
			fmt.Printf("cannot convert '%s' to an integer: %s\n", v6IncrementByFlag, err)
			os.Exit(1)
		}
		fmt.Println(iplib.IncrementIP6By(ip, z))
	},
}

func init() {
	v6RootCmd.AddCommand(v6IncrementCmd)
	v6IncrementCmd.Flags().StringVar(&v6IncrementByFlag, "by", "1", "increment address by count")

}
