package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var hexToIPCmd = &cobra.Command{
	Use:   "hextoip",
	Short: "convert hexadecimal to an IP address",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a single IP address as argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(iplib.HexStringToIP(args[0]))
	},
}

func init() {
	convertRootCmd.AddCommand(hexToIPCmd)
}
