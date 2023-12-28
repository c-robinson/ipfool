package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v4FromIntCmd = &cobra.Command{
	Use:   "int",
	Short: "integer to IPv4 address",
	Long: `
The 'v4 from int' subcommand converts an integer into an IPv4 address, where
0 is converted to 0.0.0.0 and 4294967295 becomes 255.255.255.255`,
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
