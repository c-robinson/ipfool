package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var netCmd = &cobra.Command{
	Use:   "net",
	Short: "view and manipulate ipv4 and ipv6 addresses",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a single IP address as argument")
		}
		_, _, err := iplib.ParseCIDR(args[0])
		if err != nil {
			return fmt.Errorf("could not parse network: %s", err.Error())
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		_, ipnet, _ := iplib.ParseCIDR(args[0])
		if ipnet.Version() == 4  {
			ViewIPv4Address(&ipnet)
		} else {
			ViewIPv6Address(&ipnet)
		}
	},
}

func ViewIPv4Address(ipnet *iplib.Net) {
	data := map[string]string{
		"Address": ipnet.IP.String(),
		"Netmask": ipnet.Mask.String(),
		"Network": ipnet.NetworkAddress().String(),
		"First": ipnet.FirstAddress().String(),
		"Last": ipnet.LastAddress().String(),
		"Broadcast": ipnet.BroadcastAddress().String(),
		"Count": fmt.Sprintf("%d", ipnet.Count6()),
	}

	for _, k := range []string{"Address", "Netmask", "Network", "First", "Last", "Broadcast", "Count"} {
		fmt.Printf("%-18s %-16s\n", k, data[k])
	}
}

func ViewIPv6Address(ipnet *iplib.Net) {
	data := map[string]string{
		"Address": ipnet.IP.String(),
		"Netmask": ipnet.Mask.String(),
		"First": ipnet.FirstAddress().String(),
		"Last": ipnet.LastAddress().String(),
		"Count": fmt.Sprintf("%d", ipnet.Count6()),
	}

	for _, k := range []string{"Address", "Netmask", "First", "Last", "Count"} {
		fmt.Printf("%-18s %-16s\n", k, data[k])
	}
}

func init() {
	rootCmd.AddCommand(netCmd)
}
