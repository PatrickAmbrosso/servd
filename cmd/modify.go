package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const modifyHelp = "Change service configurations (restart behavior, paths, etc.)"

var modifyCmd = &cobra.Command{
	Use:   "modify <service-name>",
	Short: modifyHelp,
	Long:  fmt.Sprintln(banner + "\n" + modifyHelp),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("modify called")
	},
}

func init() {
	rootCmd.AddCommand(modifyCmd)
}
