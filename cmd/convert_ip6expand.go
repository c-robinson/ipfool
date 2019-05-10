package cmd

import (
	"errors"
	"fmt"
	"net"

	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var ip6ExpandCmd = &cobra.Command{
	Use:   "ip6expand",
	Short: "print full IPv6 address instead of the normal, concise format",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a single IPv6 address as argument")
		}
		if ip := net.ParseIP(args[0]); ip == nil {
			return errors.New("address is not a valid IP address")
		} else if iplib.EffectiveVersion(ip) != 6 {
			return errors.New("address is not a valid IPv6 address")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ip := net.ParseIP(args[0])
		fmt.Println(iplib.ExpandIP6(ip))
	},
}

func init() {
	convertRootCmd.AddCommand(ip6ExpandCmd)
}
