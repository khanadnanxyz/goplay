package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// calcCmd represents the calc command
var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("calc called")
	},
}

func init() {
	rootCmd.AddCommand(calcCmd)
}
