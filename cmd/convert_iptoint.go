package cmd

import (
	"errors"
	"fmt"
	"net"

	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var ipToIntCmd = &cobra.Command{
	Use:   "iptoint",
	Short: "convert IPv4 or IPv6 address to integer",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a single IP address as argument")
		}
		if ip := net.ParseIP(args[0]); ip == nil {
			return errors.New("address is not a valid IP address")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ip := net.ParseIP(args[0])
		if iplib.EffectiveVersion(ip) == 4 {
			fmt.Println(iplib.IP4ToUint32(ip))
		} else {
			fmt.Println(iplib.IPToBigint(ip))
		}
	},
}

func init() {
	convertRootCmd.AddCommand(ipToIntCmd)
}
