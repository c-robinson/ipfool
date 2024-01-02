package cmd

import (
	"github.com/spf13/cobra"
)

var v4FromCmd = &cobra.Command{
	Use:                   "from",
	Short:                 "convert from in-addr.arpa, hex or integer to IPv4",
	DisableFlagsInUseLine: true,
}

func init() {
	v4RootCmd.AddCommand(v4FromCmd)
}
