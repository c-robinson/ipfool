package cmd

import (
	"fmt"
	"os"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v4CompareLongFlag bool

var v4CompareCmd = &cobra.Command{
	Use:   "compare <address1> <address2>",
	Short: "compare two IPv4 addresses",
	Long: `
The 'v4 compare' subcommand takes two IPv4 addresses as input and returns
whether the first is less than, equal to, or greater than the second. This
is somewhat useful for IPv6 but I admit it's really obvious for v4. But I
did them both anyway.

Examples:
  % ipfool v4 compare 192.168.1.1 192.168.255.1
  a < b

  % ipfool v4 compare --long 11.0.0.0 10.0.0.0
  11.0.0.0 > 10.0.0.0
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		res := []string{"<", "=", ">"}
		ipa := retrieveIPAddress(args[0], v4)
		ipb := retrieveIPAddress(args[1], v4)
		v := iplib.CompareIPs(ipa, ipb)
		if !v4CompareLongFlag {
			fmt.Printf("a %s b\n", res[v+1])
			os.Exit(0)
		}
		fmt.Printf("%s %s %s\n", ipa, res[v+1], ipb)
	},
}

func init() {
	v4RootCmd.AddCommand(v4CompareCmd)
	v4CompareCmd.Flags().BoolVarP(&v4CompareLongFlag, "long", "l", false, "display addresses in result")
}
