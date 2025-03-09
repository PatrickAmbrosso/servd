package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const restartHelp = "Restart a running service"

var restartCmd = &cobra.Command{
	Use:   "restart <service-name>",
	Short: restartHelp,
	Long:  fmt.Sprintln(banner + "\n" + restartHelp),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("restart called")
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)
}
