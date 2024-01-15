package cmd

import "github.com/spf13/cobra"

var attributesDoc = &cobra.Command{
	Use:   "attributes",
	Short: "notes on IANA forwarded, reserved, private attributes",
	Long: `
Some subcommands, most notably 'net view', report on the status of three
boolean designations contained in the IANA registries: "forwardable",
"private" and "reserved". These are copied straight from the relevant IANA
registries and their definitions are:

    Forwardable: A boolean value indicating whether a router may
      forward an IP datagram whose destination address is drawn from the
      allocated special-purpose address block between external
      interfaces.

    Private (aka "Globally Reachable"): A boolean value indicating
      whether an IP datagram whose destination address is drawn from the
      allocated special-purpose address block is forwardable beyond a
      specified administrative domain.

    Reserved(-by-Protocol): A boolean value indicating whether the
      special-purpose address block is reserved by IP, itself.  This
      value is "TRUE" if the RFC that created the special-purpose
      address block requires all compliant IP implementations to behave
      in a special way when processing packets either to or from
      addresses contained by the address block.
`,
}

func init() {
	// order matters here and the rootCmd should be last
	netRootCmd.AddCommand(attributesDoc)
	rootCmd.AddCommand(attributesDoc)
}
