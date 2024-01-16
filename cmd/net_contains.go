package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var netContainsCodeFlag bool

var netContainsCmd = &cobra.Command{
	Use:   "contains <network> <network|address>",
	Short: "check if a network encompasses another network or address",
	Long: `
'net contains' takes two arguments: a network and either a second network or
an IP address. It returns true if the second argument is wholly or partially
contained within the first, otherwise it returns false and exits code 2.

Flags:
  --code return an exit code instead of a boolean value

Examples:
  % ipfool net contains 10.0.0.0/8 192.168.255.8
  false

  % ipfool net contains 192.168.0.0/16 192.168.255.8
  true

  % ipfool net contains 2001:db8::/64 2001:db8:0:0:ffff:1:191::
  true

  % ipfool net contains 2001:db8::/64 2001:db8:0:0:1::/64
  true

  % ipfool net contains 2001:db8::/64 2001:db8:1::/64
  false

  % ipfool net contains 192.168.0.0/16 192.168.2.0/24
  true
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)

		// if the compare request is against a valid network
		_, cnet, err := iplib.ParseCIDR(args[1])
		if err == nil {

			respondToTrueFalseQuestion(ipnet.ContainsNet(cnet), netContainsCodeFlag)
		}

		caddr := net.ParseIP(args[1])
		if caddr != nil {
			respondToTrueFalseQuestion(ipnet.Contains(caddr), netContainsCodeFlag)
		}

		fmt.Println("Invalid input")
		os.Exit(1)
	},
}

func init() {
	netRootCmd.AddCommand(netContainsCmd)
	netContainsCmd.Flags().BoolVarP(&netContainsCodeFlag, "code", "x", false, "use exit code for output")
}
