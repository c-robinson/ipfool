package cmd

import (
	"fmt"
	"github.com/c-robinson/iplib"
	"github.com/spf13/cobra"
	"math/big"
)

var intToIP6Cmd = &cobra.Command{
	Use:   "inttoip6",
	Short: "convert an integer to an ip6 address",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		var MaxIPv6s = "340282366920938463463374607431768211455"
		MaxIPv6 := new(big.Int)
		MaxIPv6, _ = MaxIPv6.SetString(MaxIPv6s, 10)
		if len(args) != 1 {
			return fmt.Errorf("requires an integer value between 0 and %s", MaxIPv6s)
		}
		z := new(big.Int)
		z, ok := z.SetString(args[0], 10)
		if !ok {
			return fmt.Errorf("argument could not be converted to integer")
		}
		d := MaxIPv6.Cmp(z)
		if d == 1 {
			return fmt.Errorf("%s is greater than the IPv6 address space (%d)", args[0], MaxIPv6s)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		z := new(big.Int)
		z, _ = z.SetString(args[0], 10)
		fmt.Println(iplib.BigintToIP6(z))
	},
}

func init() {
	convertRootCmd.AddCommand(intToIP6Cmd)
}
