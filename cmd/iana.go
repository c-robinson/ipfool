package cmd

import (
	"github.com/spf13/cobra"
)

var ianaRootCmd = &cobra.Command{
	Use:   "iana",
	Short: "commands for viewing RFCs relevant to IP netblocks",
	Long: `
The 'iana' subtree holds commands for checking networks against the IANA IPv4
and IPv6 Special-Purpose Address registries. These subcommands will report on
RFCs that constrain certain netblocks or reserve them for a specific purpose.
The most well-known examples being RFC 1918 for the private IPv4 space and
RFC 3849 for the IPv6 documentation space.

These commands are just a way of viewing info scraped from the following two
URLS:
https://www.iana.org/assignments/iana-ipv4-special-registry/iana-ipv4-special-registry.xhtml
https://www.iana.org/assignments/iana-ipv6-special-registry/iana-ipv6-special-registry.xhtml
`,
}

func init() {
	rootCmd.AddCommand(ianaRootCmd)
}
