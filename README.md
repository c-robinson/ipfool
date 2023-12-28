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

Currently, you have to build it from source but that's really easy to do if
you already have a go compiler lying around. If you don't have a go compiler
lying around you can get one from [here](https://golang.org/dl/).

Now just run the following commands:

```bash
$ git clone github.com/c-robinson/ipfool
$ cd ipfool
$ go build
```

And you should have a shiny new `ipfool` binary in the current directory.

## Usage

IPFool uses a subcommand syntax similar to `git` or `docker` so if those
irritate you then I'm sorry it took me this long to reveal that and you
didn't need to follow the installation steps. If you're still with me then
here's a few quick examples of _real people just like you_ using it:


