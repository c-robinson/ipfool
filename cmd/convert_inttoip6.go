package cmd

import (
	"fmt"
	"github.com/c-robinson/iplib"
	"github.com/spf13/cobra"
	"math/big"
	"os"
)

var intToIP6Cmd = &cobra.Command{
	Use:   "inttoip6",
	Short: "integer to IPv6 address",
	Long: `
The inttoip6 subcommand converts an integer to an IPv6 address, where 0 is ::
(the all zeroes address) and 340282366920938463463374607431768211455 is
ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var MaxIPv6s = "340282366920938463463374607431768211455"
		MaxIPv6 := new(big.Int)
		MaxIPv6, _ = MaxIPv6.SetString(MaxIPv6s, 10)

		z := new(big.Int)
		z, ok := z.SetString(args[0], 10)
		if !ok {
			fmt.Println("supplied value could not be converted to integer")
			os.Exit(1)
		}
		d := MaxIPv6.Cmp(z)
		if d == -1 {
			fmt.Println("supplied value is outside the valid IPv6 address range")
			os.Exit(1)
		}

		fmt.Println(iplib.BigintToIP6(z))
	},
}

func init() {
	convertRootCmd.AddCommand(intToIP6Cmd)
}
