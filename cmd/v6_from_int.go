package cmd

import (
	"fmt"
	"os"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
	"lukechampine.com/uint128"
)

var v6FromIntCmd = &cobra.Command{
	Use:   "int <integer value>",
	Short: "IPv6 address from 128bit unsigned integer",
	Long: `
'v6 from int' converts an integer into an IPv6 address, where 0 == :: and
340282366920938463463374607431768211455 == 
ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff

Examples:
  % ipfool v6 from int 0
  ::

  % ipfool v6 from int 42540766411282592856903984951653826561
  2001:db8::1

  % ipfool v6 from int 340282366920938463463374607431768211455
  ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		z, err := uint128.FromString(args[0])
		if err != nil {
			fmt.Println("supplied value is outside the valid IPv6 address range")
			os.Exit(1)
		}
		fmt.Println(iplib.Uint128ToIP6(z))
	},
}

func init() {
	v6FromCmd.AddCommand(v6FromIntCmd)
}
