package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const startHelp = "Start a stopped service"

var startCmd = &cobra.Command{
	Use:   "start <service-name>",
	Short: startHelp,
	Long:  fmt.Sprintln(banner + "\n" + startHelp),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
