package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v4IncrementByFlag string

var v4IncrementCmd = &cobra.Command{
	Use:   "increment <address>",
	Short: "increment an IPv4 address by a given amount (default 1)",
	Long: `
The 'v4 increment' command takes an IPv4 address as input. If no arguments
are given it will increment the address by one.

Flags:
  --by <int>  increment by this amount (default 1)

Examples:
  % ipfool v4 increment 192.168.2.0
  192.168.2.1

  % ipfool v4 increment --by 256 192.168.254.1
  192.168.255.1
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"by"},
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v4)
		i, err := strconv.Atoi(v4IncrementByFlag)
		if err != nil {
			fmt.Printf("cannot convert '%s' to an integer: %s\n", v4IncrementByFlag, err)
			os.Exit(1)
		}
		fmt.Println(iplib.IncrementIP4By(ip, uint32(i)))
	},
}

func init() {
	v4RootCmd.AddCommand(v4IncrementCmd)
	v4IncrementCmd.Flags().StringVar(&v4IncrementByFlag, "by", "1", "increment address by count")

}
