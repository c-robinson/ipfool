# ipfool
IPFool is a command line tool for generating, manipulating, viewing and
generally goofing around with IP addresses and networks. Originally started
as a way to test features of the [iplib](https://github.com/c-robinson/iplib) library while that library was
being developed it turns out to honestly be pretty useful on its own if you
are in to that sort of thing. Don't get me wrong there are a lot of competing
brands in the marketplace and it's entirely possible that there are more tools
for looking at addresses than there are people who look at addresses for a
living but my feeling is I wrote the code so I might as well share it. Also
maybe this will make me hella rich somehow.

## Features
IPFool allows you to do the following things with v4 and v6 addresses:
 - convert to and from binary, hex and integer representations
 - convert to and from ARPA reverse DNS representations
 - expand v6 addresses to their full form
 - increment or decrement addresses by arbitrary amounts
 - get the delta between two addresses
 - generate IPv6 IID addresses

You can also work with lists of addresses, either parsed from a file or
provided via stdin:
 - sort addresses
 - deduplicate
 - provide a netblock and a list of used addresses and retrieve the unused ones

Technically the parser accepts lists of addresses separated by space, newline
or comma, but the parser is pretty forgiving and I regularly feed it BIND
zonefiles with no problems (because the addresses conform to the above rules
and the parser ignores things that don't look lke addresses).

But that's not all! You can also muck around with IPv4 an IPv6 netblocks:
 - view a netblock to see:
   - the first and last address in the block
   - the number of usable addresses in the block
   - the netmask and where applicable the wildcard mask
   - relevant IANA reservations and RFCs
 - subnet a netblock
 - get the netblocks supernet
 - get the previous or next netblock at the same CIDR
 - enumerate the netblock if you happen to be a masochist
 - generate a random address within the netblock
 - see if a network or address is contained within another one

## Installation
IPFool is not currently available via any major OS package managers, but
release binaries are available on the [releases](https://github.com/c-robinson/ipfool/releases)
page, in the following formats:

 - FreeBSD amd64, arm64 -- tar.gz
 - Linux x86_64, arm64 -- tar.gz, deb, rpm
 - MacOS amd64, arm64 -- zip
 - OpenBSD amd64, arm64 -- tar.gz
 - Windows amd64, arm64 - zip

## Usage
IPFool uses a subcommand syntax similar to `git` or `docker` so if those
irritate you then I'm sorry it took me this long to reveal that and you
didn't need to follow the installation steps. If you're still with me then
here's a quick overview of how it works:

There are five primary subcommands: `v4`, `v6`, `list`, `net` and `iana`. 
They're pretty self-explanatory except to note that the `list` and `net`
subcommands works for both v4 and v6. The `iana` subcommand is for looking up
IANA reservations (specifically for listing the v4 and v6 "IANA Special-Purpose
Address registries" which is where they record the netblocks that are set
aside for specific roles, like IPv4 RFC1918).

The main `v4` and `v6` subcommands have functions to find the delta between two
addresses, or to increment or decrement them by arbitrary amounts:
```bash
% ipfool v4 increment 192.168.1.1 --by 1024
192.168.5.1
% ipfool v6 increment 2001:db8::1 --by 1024
2001:db8::401
```

Both the `v4` and `v6` subcommands have `from` and `to` subcommands which
provide for conversion between various representations of addresses:
```bash
% ipfool v4 to arpa 192.168.1.1               
1.1.168.192.in-addr.arpa
% ipfool v4 from arpa 1.1.168.192.in-addr.arpa
192.168.1.1
% ipfool v6 to arpa 2001:db8::1
1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa
% ipfool v6 from arpa 1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa
2001:db8::1
```

There are also options to convert to and from binary and integer
representations. IPv4 can also be converted to hexadecimal; while IPv6 can
be expanded to its full form.

Finally, the `v6` subcommand has a function to generate an EUI64-compliant
IID address, or an RFC7217-compliant "semantically opaque" address:
```bash
% ipfool v6 iid 2001:db8:: 01:23:45:67:89:ab
2001:db8::123:45ff:fe67:89ab
% ipfool v6 iid --counter 1 --secret P00P570Rm 2001:db8:: 01:23:45:67:89:ab
2001:db8::7da9:ddbb:107a:8c07
```

The `list` subcommand has functions to sort and deduplicate lists of IPs,
and can often tease them out of files that have a bunch of other garbage in
them so long as the addresses are surrounded by either whitespace or commas.
```bash
% ipfool list free 192.168.235.0/24 --file /tmp/db.my-zonefile --summary=only
Network: 192.168.235.0/24
  Size: 254,  Used: 110 (43.31%)  Free: 145
```

The `net` subcommand has functions to work with v4 or v6 netblocks:
```bash
% ipfool ipfool.go net subnet --cidr 24 192.168.0.0/22
192.168.0.0/24
192.168.1.0/24
192.168.2.0/24
192.168.3.0/24
% ipfool net view 2001:db8::/64                                            
Address            2001:db8::      
Netmask            ffff:ffff:ffff:ffff:0000:0000:0000:0000
First              2001:0db8:0000:0000:0000:0000:0000:0000
Last               2001:0db8:0000:0000:ffff:ffff:ffff:ffff
Count              18446744073709551616
Registered in: RFC3849
Network may not be forwarded, is private, is not reserved
````

For IPv6 addresses `ipfool` supports the (entirely optional) concept of a
"hostmask", which masks bits from the right-hand side of the netblock in
the same way that netmask does from the left. Hostmasks are assigned by adding
them to the CIDR-notated netblock after a colon, for example:
`2001:db8::/64:48`; the additional mask allows you to further constrain many
of `ipfool net`'s functions in the v6 context:
```bash
% ipfool net view 2001:db8::/64:48                                           
Address            2001:db8::      
Netmask            ffff:ffff:ffff:ffff:0000:0000:0000:0000
Hostmask           0000:0000:0000:0000:0000:ffff:ffff:ffff
First              2001:0db8:0000:0000:0000:0000:0000:0000
Last               2001:0db8:0000:0000:ffff:0000:0000:0000
Count              65536           
Registered in: RFC3849
Network may not be forwarded, is private, is not reserved
```

Note that there are now only 65536 addressable entities in the netblock,
where without a hostmask there are 18446744073709551616. Peep how this
crazy vibe changes incrementing, as one example:
```bash
% ipfool net increment 2001:db8::/64 2001:db8:: 
2001:db8::1
% ipfool net increment 2001:db8::/64:48 2001:db8:: 
2001:db8:0:0:1::
```
Anyhow, those are the basics.  Documentation is available for any subcommand
at any  level via the `-h` flag.

### Building
Building IPFool is pretty easy if you're familiar with Go and have the
compiler toolchain floating around. Just run the following commands:

```bash
$ git clone github.com/c-robinson/ipfool
$ cd ipfool
$ go mod tidy
$ go build
```

It is also possible to build the entire IPFool distribution set by
using the [goreleaser](https://goreleaser.com) tool, with the caveat
that official OSX and Windows builds are signed so the versions you
will produce are not suitable for distribution. Assuming you are in the
root of the repository and have the goreleaser tool installed you can
just run:

```bash
% goreleaser release --snapshot --clean
```

Which is what I do when I'm testing junk. It will create a directory
`dist` containing the binaries for all supported platforms.