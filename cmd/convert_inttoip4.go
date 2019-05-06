package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var intToIP4Cmd = &cobra.Command{
	Use:   "inttoip4",
	Short: "convert an integer to an ip4 address",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("requires an integer value between 0 and %d", iplib.MaxIPv4)
		}
		i, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("argument could not be converted to integer: %s", err.Error())
		}
		if i > iplib.MaxIPv4 {
			return fmt.Errorf("%d is greater than the IPv4 address space (%d)", i, iplib.MaxIPv4)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		i, _ := strconv.Atoi(args[0])
		fmt.Println(iplib.Uint32ToIP4(uint32(i)))
	},
}

func init() {
	convertRootCmd.AddCommand(intToIP4Cmd)
}
