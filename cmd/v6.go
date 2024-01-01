package cmd

import "github.com/spf13/cobra"

var v6RootCmd = &cobra.Command{
	Use:   "v6",
	Short: "commands for working with IPv6 addresses",
	Long: `
The 'v6' subtree holds commands for working with IPv6 addresses. This
includes:

 * converting between IPv6 and in-addr.arpa, binary, or integer
   representations
 * expanding IPv6 addresses so all leading zeroes are present
 * calculating the distance between two IPv6 addresses
 * incrementing or decrementing an IPv6 address

To work with IPv4 networks, use the 'net' subtree.
	`,
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(v6RootCmd)
}
