package cmd

import (
	"fmt"
	"os"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v6CompareLongFlag bool

var v6CompareCmd = &cobra.Command{
	Use:   "compare <address1> <address2>",
	Short: "compare two IPv6 addresses",
	Long: `
'v6 compare' tests two addresses for equality.

Flags:
  -l, --long   display addresses in result

Examples:
  % ipfool v6 compare 2001:db8::1 2001:db8::ffff
  a < b

  % v6 compare --long 2001:db8::1 2001:db8:1::
  2001:db8::1 < 2001:db8:1::
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		res := []string{"<", "=", ">"}
		ipa := retrieveIPAddress(args[0], v6)
		ipb := retrieveIPAddress(args[1], v6)
		v := iplib.CompareIPs(ipa, ipb)
		if !v6CompareLongFlag {
			fmt.Printf("a %s b\n", res[v+1])
			os.Exit(0)
		}
		fmt.Printf("%s %s %s\n", ipa, res[v+1], ipb)
	},
}

func init() {
	v6RootCmd.AddCommand(v6CompareCmd)
	v6CompareCmd.Flags().BoolVarP(&v6CompareLongFlag, "long", "l", false, "display addresses in result")
}
