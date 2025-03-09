package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const monitorHelp = "Track service status for changes"

var monitorCmd = &cobra.Command{
	Use:   "monitor <service-name|OPTIONAL>",
	Short: monitorHelp,
	Long:  fmt.Sprintln(banner + "\n" + monitorHelp),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("monitor called")
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
}
