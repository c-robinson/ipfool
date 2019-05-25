package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var netEnumerateCmd = &cobra.Command{
	Use:   "enumerate",
	Short: "print all IPs in the subnet (caveat emptor)",
	Long:  "",
	DisableFlagsInUseLine: true,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ipnet := retrieveIPNetwork(args[0], v46)
		ip := ipnet.FirstAddress()
		for {
			fmt.Println(ip.String())
			ip, err = ipnet.NextIP(ip)
			if err != nil {
				break
			}
		}
	},
}

func init() {
	netRootCmd.AddCommand(netEnumerateCmd)
}
