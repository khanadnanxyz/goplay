package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// mulCmd represents the mul command
var mulCmd = &cobra.Command{
	Use:   "mul",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mul called")
		mul(args)
	},
}

func init() {
	calcCmd.AddCommand(mulCmd)
}

func mul(args []string)  {
	res := 1
	count := 0
	for _, item := range args {
		count = count + 1
		fmt.Printf("item no %d: %s \n" , count, item)
		i, err := strconv.Atoi(item)
		res = res * i
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(res)
}