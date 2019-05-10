package cmd

import (
	"errors"
	"fmt"
	"net"

	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var ipToARPACmd = &cobra.Command{
	Use:   "iptoarpa",
	Short: "IPv4 or IPv6 address to DNS ARPA-domain PTR",
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
		fmt.Println(iplib.IPToARPA(ip))
	},
}

func init() {
	convertRootCmd.AddCommand(ipToARPACmd)
}
