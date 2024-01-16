package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v4FromIntCmd = &cobra.Command{
	Use:   "int <integer value>",
	Short: "IPv4 address from 32bit unsigned integer",
	Long: `
'v4 from int' converts an integer into an IPv4 address where 0 == 0.0.0.0
and 4294967295 == 255.255.255.255

Examples:
  % ipfool v4 from int 3232235777
  192.168.1.1
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("supplied value is outside the valid IPv4 address range")
			os.Exit(1)
		}
		fmt.Println(iplib.Uint32ToIP4(uint32(i)))
	},
}

func init() {
	v4FromCmd.AddCommand(v4FromIntCmd)
}
