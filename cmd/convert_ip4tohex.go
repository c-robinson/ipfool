package cmd

import (
	"errors"
	"fmt"
	"net"

	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var ip4ToHexCmd = &cobra.Command{
	Use:   "ip4tohex",
	Short: "dotted-decimal address to hexadecimal",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a single IPv4 address as argument")
		}
		if ip := net.ParseIP(args[0]); ip == nil {
			return errors.New("address is not a valid IPv4 address")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ip := net.ParseIP(args[0])
		fmt.Println(iplib.IPToHexString(ip))
	},
}

func init() {
	convertRootCmd.AddCommand(ip4ToHexCmd)
}
