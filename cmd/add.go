package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Do some add task",
	Run: func(cmd *cobra.Command, args []string) {
		add(args)
	},
}

func init() {
	calcCmd.AddCommand(addCmd)
}

func add(args []string)  {
	sum := 0
	count := 0
	for _, item := range args {
		count = count + 1
		fmt.Printf("item no %d: %s \n" , count, item)
		i, err := strconv.Atoi(item)
		sum = sum + i
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(sum)
}
