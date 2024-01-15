package cmd

import (
	"github.com/spf13/cobra"
)

var aboutIIDDoc = &cobra.Command{
	Use:   "about-iid",
	Short: "notes on IPv6 EUI64 Interface Identifiers",
	Long: `
The 'v6 make-iid' subcommand expects two arguments: an IPv6 address and a
hardware address (typically the target interface's MAC address). Given no
other inputs it will generate an EUI64-style IPv6 IID scoped for local-subnet
use ONLY, this is because and IID generated in this way leaks personal
information about the addresses owner (specifically it leaks the MAC address
of the host, which can be problematic for laptops or mobile devices). Note
that the provided MAC must be either 48bits or 64bits. If you provide a 48bit
MAC address, it will be padded with 0xfffe in the middle to make it 64bits.

There are three additional flags to 'make-iid' that can be used to generate a
"semantically opaque" IID as described in RFC 7217:

  --secret   REQUIRED. Should be given a secret key, theoretically from a
             secure source like an ssh private key stored in a file only you
             can read in a directory only you can access. In practice, you
             can give it any text you want.

  --neighbor OPTIONAL. To comply with the spec, this would be the MAC
             address of the next-hop router. In practice, you can give it
             any text you want, or just don't use it at all. It's supposed
             to improve the randomness of the result.

  --count    OPTIONAL. Again to comply with the spec, this would be a
             monotonically incrementing number stored someplace, retrieved
             for use and incremented each time. Again the goal is to add
             entropy to the result.

This particular implementation uses SHA256 as the hashing function. There
is theoretically one more toggleable element defined in the RFC, the "scope",
but it's very ambiguously defined so this implementation doesn't support it.
The explanation is boring but if you're curious here's the relevant comment
from the code:

  // Scope describes the availability of an IPv6 IID and determines how IID-
  // generating functions treat the 7th bit in the 9th octet of the address
  // (the 'X' bit in the EUI-64 format, or the 'u' bit in RFC4291)
  //
  // NOTE: there is some ambiguity to the RFC here. Most discussions I've seen
  // on the topic say that the bit should _always_ be inverted, but the RFC
  // reads like the IPv6 EUI64 format uses the _inverse signal_ from the IEEE
  // EUI64 format; so where the IEEE uses 0 for global scoping, the IPv6 IID
  // should use 1. This module punts on the question and provides for all
  // interpretations via the scope parameter but recommends passing an explicit
  // ScopeGlobal or ScopeLocal

I promise boring and I deliver! Anyhow as mentioned above we always pass 0,
which is probably "local", but I'm not aware of anything that actually uses
the bit; so in practice the only effect that it is harder to coerce the MAC
address out of the IID since the disposition of that bit is ambiguous.
`,
}

func init() {
	// order matters here and the rootCmd should be last
	v6RootCmd.AddCommand(aboutIIDDoc)
	rootCmd.AddCommand(aboutIIDDoc)
}
