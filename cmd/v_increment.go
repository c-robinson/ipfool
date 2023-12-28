package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
	"lukechampine.com/uint128"
)

var incBy string // string so it can be sent to big.Int.SetString()

var v46IncrementCmd = &cobra.Command{
	Use:   "increment",
	Short: "increment an IP address by a given amount (default 1)",
	Long: `
The 'increment' command takes an IP address as input. If no arguments are
given it will increment the address by one. the --by flag can be used to
specify an increment count larger than 1 (or 1, or 0 frankly).`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"by"},
	Run: func(cmd *cobra.Command, args []string) {
		switch parent := cmd.Parent().Name(); parent {
		case "v4":
			ip := retrieveIPAddress(args[0], v4)
			i, err := strconv.Atoi(decBy)
			if err != nil {
				fmt.Printf("cannot convert '%s' to an integer: %s\n", decBy, err)
				os.Exit(1)
			}
			fmt.Println(iplib.DecrementIPBy(ip, uint32(i)))

		case "v6":
			ip := retrieveIPAddress(args[0], v6)
			z, err := uint128.FromString(decBy)
			if err != nil {
				fmt.Printf("cannot convert '%s' to an integer: %s\n", decBy, err)
				os.Exit(1)
			}
			fmt.Println(iplib.DecrementIP6By(ip, z))
		}
	},
}

func init() {
	var v4IncrementCmd = *v46IncrementCmd
	var v6IncrementCmd = *v46IncrementCmd
	v4RootCmd.AddCommand(&v4IncrementCmd)
	v6RootCmd.AddCommand(&v6IncrementCmd)
	v4IncrementCmd.Flags().StringVar(&incBy, "by", "1", "increment address by count")
	v6IncrementCmd.Flags().StringVar(&incBy, "by", "1", "increment address by count")

}
