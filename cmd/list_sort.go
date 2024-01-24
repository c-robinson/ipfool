package cmd

import (
	"fmt"
	"sort"

	"github.com/c-robinson/iplib/v2"
	"github.com/spf13/cobra"
)

var listSortUniqFlag bool

var listSortCmd = &cobra.Command{
	Use:   "sort <list...>",
	Short: "sort a list of IP addresses",
	Long: `
'list sort' takes a list of IP addresses via file or stdin and sorts it.

Examples:
  % echo 172.25.11.0,192.168.32.15,10.0.1.256 | ipfool list sort -
  10.0.1.256
  172.25.11.0
  192.168.32.15
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.MaximumNArgs(1),
	ValidArgs:             []string{"file"},
	RunE: func(cmd *cobra.Command, args []string) error {
		iplist, isStdin, err := retrieveIPList(args, cmd.InOrStdin())
		if err != nil {
			return err
		}
		if len(iplist) == 0 {
			return nil
		}
		sort.Sort(iplib.ByIP(iplist))
		if listSortUniqFlag {
			iplist = ipListUniquer(iplist)
		}
		if isStdin {
			// this is to provide some visual separation between the input and
			// output when using stdin
			fmt.Println("\n---")
		}
		for _, ip := range iplist {
			fmt.Println(ip)
		}
		return nil
	},
}

func init() {
	listRootCmd.AddCommand(listSortCmd)
	listSortCmd.Flags().BoolVarP(&listSortUniqFlag, "uniq", "u", false, "prune duplicates")
}
