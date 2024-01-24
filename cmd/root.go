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
Convert IPv4 addresses between formats, increment or decrement them by a given
amount, or calculate the delta between two addresses. Conversions include
in-addr.arpa, binary, hexadecimal and integer representations.

v6
Convert IPv6 addresses between formats, increment or decrement them by a given
amount, or calculate the delta between two addresses. Conversions include
ip6.arpa, binary and integer representations, as well as "expanding" the
address. Also available is the 'iid' subcommand, which generates an RFC7217
compliant "semantically opaque" IPv6 address from bits and bobs you supply.

list
Work on lists of IP addresses. Sort them, prune duplicates, return their
inverse within a netblock (e.g. find the free addresses).

net
Work with IPv4 and IPv6 networks. Increment or decrement within a netblock,
generate a random IP from the netblock or enumerate all or part of it.
Calculate sub- and supernets as well as adjacent neighbors. View details of
the netblock.

iana
Query the Special Purpose Registries by RFC, IP or netblock to find constraints
such as RFC1918 -> 192.168.0.0/16, or RFC3849 -> 2001:DB8::/32.

For v6 addresses within the net subtree, this tool introduces the concept
of a "hostmask", which is an optional secondary mask that can be applied to
a netblock to constrain it's host bits. The mask is optional and more info
can be found via 'ipfool help hostmask'.
`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

// Execute runs rootCmd and needs normal printer in case onIntialize fails
// and no printer is available
func Execute(version string) {
	Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("while executing command: %s\n", err)
		os.Exit(1)
	}
}
