package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ipfool",
	Short: "IP Fool is a utility for monkeying with IP addresses and subnets",
	Long: `
IP Fool is a feature-rich utility for viewing and manipulating IPv4 and IPv6
addresses and subnets. It uses a "git style" syntax, grouping functions under
sub-commands to hopefully make the tool more coherent and easier to navigate.

v4
This subtree holds commands for converting IPv4 addresses between formats,
incrementing or decrementing them by a given amount, or calculating the
delta between two addresses. Conversions include in-addr.arpa, binary,
hexadecimal and integer representations.

v6
This subtree holds commands for converting IPv6 addresses between formats,
incrementing or decrementing them by a given amount, or calculating the
delta between two addresses. Conversions include in-addr.arpa, binary and
integer representations, as well as "expanding" the address. Also available
is the 'iid' subcommand, which generates an RFC 7217-compliant "semantically
opaque" IPv6 address from bits and bobs you give it.

net
This subtree holds commands for working with both IPv4 and IPv6 networks.
The 'net view' subcommand is probably the most generally interesting as it
provides a summary of the subnet (size, first and last address, etc). Other
commands provide tools for calculating the previous or next neighbor of a
block, super- or sub-netting the block to a given mask length, generating
random addresses for a block or enumerating part or all of it.

Further help can be found for each subcommand via:
	ipfool <subcommand> --help
`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

// Execute runs rootCmd and needs normal printer in case onIntialize fails
// and no printer is available
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("while executing command: %s\n", err)
		os.Exit(1)
	}
}
