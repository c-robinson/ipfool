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
flag. 

For IPv6 addresses ipfool supports the notion of a "hostmask", which
is described in more detail via 'ipfool help hostmask'. This option is
transparent, meaning that if you don't invoke the command with a
hostmask you won't get any information about it. The brief explanation
is that a hostmask masks bits from the right side of an address in the
same way that netmask does from the left and can be used to constrain
the following commands within the 'net' subtree: 'count', increment',
'decrement' and 'enumerate'. Hostmasks are specified by appending the mask
size to the address with a ':', for example '2001:db8::/56:64' defines
a network with a 56bit netmask and a 64bit hostmask:

  % ipfool net view 2001:db8::/56:64
  Address            2001:db8::
  Netmask            ffff:ffff:ffff:ff00:0000:0000:0000:0000
  Hostmask           0000:0000:0000:0000:ffff:ffff:ffff:ffff
  First              2001:0db8:0000:0000:0000:0000:0000:0000
  Last               2001:0db8:0000:00ff:0000:0000:0000:0000
  Count              256
  Registered in: RFC3849
  Network may not be forwarded, is private, is not reserved
`,
}

func init() {
	rootCmd.AddCommand(netRootCmd)
}
