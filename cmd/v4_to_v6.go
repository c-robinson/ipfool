package cmd

import (
	"fmt"
	"strings"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v4Tov6Cmd = &cobra.Command{
	Use:   "v6 <address>",
	Short: "IPv4 address to 'v4 mapped IPv6' address",
	Long: `
'v4 to v6' converts a given IPv4 address to a 'v4 mapped v6' address as defined
in RFC 4291 section 2.5.5.2. From the relevant section:

  |                80 bits               | 16 |      32 bits        |
  +--------------------------------------+--------------------------+
  |0000..............................0000|FFFF|    IPv4 address     |
  +--------------------------------------+----+---------------------+

Examples:
  % ipfool v4 to v6 192.168.1.1                    
  ::ffff:c0a8:0101
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := retrieveIPAddress(args[0], v4)
		v6 := iplib.ExpandIP6(ip)

		fmt.Println(strings.Replace(v6, "0000:0000:0000:0000:0000", ":", 1))
	},
}

func init() {
	v4ToCmd.AddCommand(v4Tov6Cmd)
}
