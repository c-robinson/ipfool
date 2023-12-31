package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var ncCode bool

var netContainsCmd = &cobra.Command{
	Use:   "contains",
	Short: "does the given network contain the provided network or address",
	Long: `
The 'net contains' subcommand takes two arguments: a network and either a
second network or an IP address. It returns true if the second argument is
wholly or partially contained within the first. 

If the --code flag is provided then output is returned as an exit code (0 for
true, 1 for false). This is a little problematic since a non-zero exit code
is overloaded for "doesn't contain" as well as input errors but technically
the word "pickle" is not contained within the 2001:db8::/64 address space so
it kind of works out.`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ipnet := retrieveIPNetwork(args[0], v46)

		// if the compare request is against a valid network
		_, cnet, err := iplib.ParseCIDR(args[1])
		if err == nil {
			respondToTrueFalseQuestion(ipnet.ContainsNet(cnet), ncCode)
		}

		caddr := net.ParseIP(args[1])
		if caddr != nil {
			respondToTrueFalseQuestion(ipnet.Contains(caddr), ncCode)
		}

		fmt.Println("Invalid input")
		os.Exit(1)
	},
}

func init() {
	netRootCmd.AddCommand(netContainsCmd)
	netContainsCmd.Flags().BoolVarP(&ncCode, "code", "x", false, "use exit code for output")
}

func respondToTrueFalseQuestion(result, exitCodeOnly bool) {
	if !exitCodeOnly {
		fmt.Printf("%t\n", result)
	}
	if result {
		os.Exit(0)
	}
	os.Exit(1)
}
