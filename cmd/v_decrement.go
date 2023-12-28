package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
	"lukechampine.com/uint128"
)

var decBy string

var v46DecrementCmd = &cobra.Command{
	Use:   "decrement",
	Short: "decrement an IP address by a given amount (default 1)",
	Long: `
The 'decrement' command takes an IP address as input. If no arguments are
given it will decrement the address by one. the --by flag can be used to
specify a decrement count larger than 1 (or 1, or 0 frankly).`,
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
	var v4DecrementCmd = *v46DecrementCmd
	v4RootCmd.AddCommand(&v4DecrementCmd)
	v4DecrementCmd.Flags().StringVar(&decBy, "by", "1", "decrement address by count")

	var v6DecrementCmd = *v46DecrementCmd
	v6RootCmd.AddCommand(&v6DecrementCmd)
	v6DecrementCmd.Flags().StringVar(&decBy, "by", "1", "decrement address by count")

}
