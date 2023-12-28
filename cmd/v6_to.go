package cmd

import (
	"github.com/spf13/cobra"
)

var v6ToCmd = &cobra.Command{
	Use:                   "to",
	Short:                 "convert from IPv6 to in-addr.arpa, binary or integer",
	DisableFlagsInUseLine: true,
}

func init() {
	v6RootCmd.AddCommand(v6ToCmd)
}
