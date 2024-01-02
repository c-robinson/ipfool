package cmd

import (
	"fmt"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v6ToIntCmd = &cobra.Command{
	Use:   "int <address>",
	Short: "IPv6 address to 128bit unsigned integer",
	Long: `
The 'v6 to int' subcommand converts a given IPv6 address into an unsigned
128bit integer where the first address (::) is indexed to 0. Most of the IP
math in this library is performed by converting addresses to integers and
performing arithmatic operations on them in that format, since it's much
simpler and way faster than doing bit operations. Why you'd want to do it
on the command line is an exercise fo the reader but I'm providing this
anyway under an "every part of the buffalo" clause.

Examples:
  % ipfool v6 to int ::
  0

  % ipfool v6 to int 
  42540766411282592856903984951653826561

  % ipfool v6 to int ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff
  340282366920938463463374607431768211455
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v6)
		fmt.Println(iplib.IP6ToUint128(ip))
	},
}

func init() {
	v6ToCmd.AddCommand(v6ToIntCmd)
}
