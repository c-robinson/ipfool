package cmd

import (
	"github.com/spf13/cobra"
)

var netRootCmd = &cobra.Command{
	Use:   "net",
	Short: "commands for viewing and manipulating IP netblocks",
	Long: `
The 'net' subtree holds commands for viewing IP netblocks as well as
experimenting with subnets and supernets.

If you work with IPv6 networks it might be useful to read this brief intro
to a thing I call a "hostmask" as it might make your life easier. A hostmask
functions like a netmask, but instead of masking bits from the left it does
so from the right; this allows you to constrain actions like 'enumerate',
'next' and 'previous' to a narrower range than using netmask alone.
`,
}

func init() {
	rootCmd.AddCommand(netRootCmd)
}
