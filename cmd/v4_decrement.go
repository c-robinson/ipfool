package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v4DecrementByFlag string

var v4DecrementCmd = &cobra.Command{
	Use:   "decrement <address>",
	Short: "decrement an IPv4 address by a given amount (default 1)",
	Long: `
The 'decrement' command takes an IP address as input. If no arguments are
given it will decrement the address by one. the --by flag can be used to
specify a decrement count larger than 1 (or 1, or 0 frankly).

Examples:
  % ipfool v4 decrement 192.168.2.1
  192.168.2.0

  % ipfool v4 decrement --by 256 192.168.2.1  
  192.168.1.1
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	ValidArgs:             []string{"by"},
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v4)
		i, err := strconv.Atoi(v4DecrementByFlag)
		if err != nil {
			fmt.Printf("cannot convert '%s' to an integer: %s\n", v4DecrementByFlag, err)
			os.Exit(1)
		}
		fmt.Println(iplib.DecrementIPBy(ip, uint32(i)))
	},
}

func init() {
	v4RootCmd.AddCommand(v4DecrementCmd)
	v4DecrementCmd.Flags().StringVar(&v4DecrementByFlag, "by", "1", "decrement IPv4 address by count")
}
