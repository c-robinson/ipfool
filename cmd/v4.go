package cmd

import "github.com/spf13/cobra"

var v4RootCmd = &cobra.Command{
	Use:   "v4",
	Short: "commands for working with IPv4 addresses",
	Long: `
The 'v4' subtree holds commands for working with IPv4 addresses. This
includes:

 * converting between IPv4 and in-addr.arpa, binary, hexadecimal or
   integer representations
 * calculating the distance between two IPv4 addresses
 * incrementing or decrementing and IPv4 address

To work with IPv4 networks, use the 'net' subtree.
	`,
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(v4RootCmd)
}
