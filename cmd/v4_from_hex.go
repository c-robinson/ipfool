package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v4FromHexCmd = &cobra.Command{
	Use:   "hex",
	Short: "hexadecimal to IPv4 address",
	Long: `
The 'v4 from hex' subcommand converts a hexadecimal value to an IPv4 address,
for example 0xc0a80101 becomes 192.168.1.1

The 0x prefix is optional but harmless, knock yourself out.`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hex := strings.TrimLeft(args[0], "0x")
		ip := iplib.HexStringToIP(hex)
		if ip == nil {
			fmt.Printf("invalid hex value: %s\n", hex)
			os.Exit(1)
		}
		fmt.Println(ip)
	},
}

func init() {
	v4FromCmd.AddCommand(v4FromHexCmd)
}
