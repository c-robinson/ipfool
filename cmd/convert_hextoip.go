package cmd

import (
	"fmt"
	"github.com/c-robinson/iplib"
	"github.com/spf13/cobra"
	"strings"
)

var hexToIPCmd = &cobra.Command{
	Use:   "hextoip",
	Short: "hexadecimal to IP address",
	Long: `
The hextoip subcommand converts a hexadecimal value to an IP address. This is
meaningless to IPv6 where the default representation is already hexadecimal
so is mostly only valuable for converting into IPv4; for example 0xc0a80101
is 192.168.1.1`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hex := strings.TrimLeft(args[0], "0x")
		fmt.Println(iplib.HexStringToIP(hex))
	},
}

func init() {
	convertRootCmd.AddCommand(hexToIPCmd)
}
