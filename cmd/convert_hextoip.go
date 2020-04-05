package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var hexToIPCmd = &cobra.Command{
	Use:                   "hextoip",
	Short:                 "hexadecimal to IP address",
	Long:                  "",
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(iplib.HexStringToIP(args[0]))
	},
}

func init() {
	convertRootCmd.AddCommand(hexToIPCmd)
}
