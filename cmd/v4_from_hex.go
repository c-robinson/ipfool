package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var v4FromHexCmd = &cobra.Command{
	Use:   "hex <hexadecimal value>",
	Short: "IPv4 address from hexadecimal",
	Long: `
The 'v4 from hex' subcommand converts a hexadecimal value to an IPv4 address.
The traditional '0x' prefix is optional but knock yourself out.

Examples:
  % ipfool v4 from hex 0xc0a80101
  192.168.1.1
`,
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
