package cmd

import (
	"github.com/spf13/cobra"
)

var netRootCmd = &cobra.Command{
	Use:   "net",
	Short: "commands for viewing and manipulating IP netblocks",
	Long: `
The 'net' subtree holds commands for viewing IP netblocks as well as
experimenting with subnets and supernets. The most generally useful
subcommand might be 'net view' which displays information about the
provided netblock. Note that most commands under this subtree can be
told to report their results in the view format by using the --view
flag. For example:

  % ipfool net nextnet --view 192.168.0.0/24
  Original           192.168.0.0/24
  Address            192.168.0.0
  Netmask            255.255.255.0
  Network            192.168.0.0
  First              192.168.0.1
  Last               192.168.0.254
  Wildcard           000000ff
  Broadcast          192.168.0.255
  Count              254
  Registered in: RFC1918
  Network may be forwarded, is private, is not reserved

  Next adjacent      192.168.1.0/24
  Address            192.168.1.0
  Netmask            255.255.255.0
  Network            192.168.1.0
  First              192.168.1.1
  Last               192.168.1.254
  Wildcard           000000ff
  Broadcast          192.168.1.255
  Count              254
  Registered in: RFC1918
  Network may be forwarded, is private, is not reserved
`,
}

func init() {
	rootCmd.AddCommand(netRootCmd)
}
