package cmd

import "github.com/spf13/cobra"

var listFileFlag string

var listRootCmd = &cobra.Command{
	Use:   "list",
	Short: "commands for working with lists of IP addresses",
	Long: `
The 'list' subtree holds various commands for working with lists of IP
addresses. This includes:

 * sorting a list of addresses, including pruning duplicates
 * returning a list of addresses that are missing from a supplied list,
   including calculating how much of a subnet is used/free

commands in this subtree can take lists in one of two ways:
 * via the --file flag, which takes a filename as an argument
 * via stdin, with a single dash (-) supplied as the final argument

A note on input: the "hard" rule is addresses may be comma, space or newline
separated, but the parser is fairly forgiving beyond that. For example I
regularly feed it entire BIND zone files and it happily parses the addresses
out of them.
`,
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(listRootCmd)
	listRootCmd.PersistentFlags().StringVarP(&listFileFlag, "file", "f", "", "filename to use")
}
