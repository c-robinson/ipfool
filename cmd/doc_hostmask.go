package cmd

import "github.com/spf13/cobra"

var hostmaskDoc = &cobra.Command{
	Use:   "hostmask",
	Short: "notes on IPv6 hostmasks",
	Long: `
HostMask is a mask that can be applied to IPv6 addresses to mask out bits
from the right side of the address instead of the left (which is the
purview of a netmask), the intended use-case is for situations where there
is a desire to reserve a portion of the address for some other purpose and
only allow iplib to manage the remainder. A concrete example would be 
IPv6 Interface Identifiers as described in RFC4291, RFC4941 or RFC7217 in 
which the final 64bits of the address are used to construct a unique host
identifier and the allocator only has control of the first 64bits. So the
next IP from 2001:db8:1234:5678:: would be 2001:db8:1234:5679 instead of 
2001:db8:1234:5678::1. To use a hostmask, append the mask size after the
CIDR value with a ':', for example '2001:db8::/56:60' defines a network with
a 56bit netmask and a 60bit hostmask.

Here is a sample view of an IPv6 netblock without a hostmask:

    % ipfool net view 2001:db8::/56    
    Address            2001:db8::      
    Netmask            ffff:ffff:ffff:ff00:0000:0000:0000:0000
    First              2001:0db8:0000:0000:0000:0000:0000:0000
    Last               2001:0db8:0000:00ff:ffff:ffff:ffff:ffff
    Count              4722366482869645213696
    Registered in: RFC3849
    Network may not be forwarded, is private, is not reserved

This creates a block with 4.7 sextillion usable addresses. Below is he same
block with a hostmask of 60, created as described above by appending ':60'.
The mask is applied from the rightmost byte, leaving 12 unmasked bits for a
total of 4096 allocatable addresses:

    % ipfool net view 2001:db8::/56:60
    Address            2001:db8::      
    Netmask            ffff:ffff:ffff:ff00:0000:0000:0000:0000
    Hostmask           0000:0000:0000:0000:f0ff:ffff:ffff:ffff
    First              2001:0db8:0000:0000:0000:0000:0000:0000
    Last               2001:0db8:0000:00ff:0f00:0000:0000:0000
    Count              4096            
    Registered in: RFC3849
    Network may not be forwarded, is private, is not reserved

In the first example the second IP address of the netblock is 2001:db8::1,
in the second example it is 2001:db8:0:1::

One important note: even though bytes are filled in from the right the bits
within those bytes are still blocked out left-to-right, so that address
incrementing/decrementing makes sense to the end user, as shown here:

    BINARY      Base16  Base10  Example Max16  Max10
    0000 0000     0x00       0      /56  0xFF    255
    1000 0000     0x80     128      /57  0x7F    127
    1100 0000     0xC0     192      /58  0x3F     63
    1110 0000     0xE0     224      /59  0x1F     31
    1111 0000     0xF0     240      /60  0x0F     15
    1111 1000     0xF8     248      /61  0x07      7
    1111 1100     0xFC     252      /62  0x03      3
    1111 1110     0xFE     254      /63  0x01      1

A hostmask of /1 will block out the left-most bit of the 16th byte
while a /8 will block the entire 16th byte.

To reiterate: I am pretty sure this is my own construction, at least I've
never seen a formal proposal for it. As a result I've taken great pains to
make sure you don't ever have to use it. It might be a very dumb idea I am
having and I wouldn't want to inflict it on this tool that I think is
otherwise quite useful. If you never use the :<hostmask> syntax you never
even have to know the feature exists.
`,
}

func init() {
	// order matters here and the rootCmd should be last
	netRootCmd.AddCommand(hostmaskDoc)
	rootCmd.AddCommand(hostmaskDoc)
}
