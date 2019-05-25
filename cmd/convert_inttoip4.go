package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/c-robinson/iplib"
)

var intToIP4Cmd = &cobra.Command{
	Use:   "inttoip4",
	Short: "integer to IPv4 address",
	Long:  "",
	DisableFlagsInUseLine: true,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("supplied value is outside the valid IPv4 address range")
			os.Exit(1)
		}
		fmt.Println(iplib.Uint32ToIP4(uint32(i)))
	},
}

func init() {
	convertRootCmd.AddCommand(intToIP4Cmd)
}
