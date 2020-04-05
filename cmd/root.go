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
	Long: `IP Fool is a feature-rich utility for viewing and manipulating IPv4 and IPv6
addresses and subnets. It uses a "git style" syntax, grouping functions under
sub-commands to hopefully make the tool more coherent and easier to navigate.

Convert
The "convert" subcommand contains functions for printing addresses in various
formats, for example viewing an IP address by its binary, hexidecimal or
integer value, or printing its ARPA reverse-DNS entry (which can save a lot
of typing in the v6 case). The "convert iid" subcommand will generate an RFC
7217-compliant "semantically opaque" IPv6 address.

Difference
The "difference" subcommand contains functions for calculating the distance
between two addresses or for incrementing/decrementing a supplied address by
a given amount.

Net
The "net" subcommand is for working with subnets. "net view" provides a
summary of the subnet (size, first and last address, etc) as well as
reporting on any IANA reservations which may apply to all or part of the
it and what RFC covers them. The rest of the commands provide tools for
enumerating that block into discrete addresses (which may take infinitely
long for IPv6), carving the block into smaller subnets, or retrieving the
supernet of the given block.
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
