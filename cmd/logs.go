package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const logsHelp = "View logs from the managed services"

var logsCmd = &cobra.Command{
	Use:   "logs <service-name>",
	Short: logsHelp,
	Long:  fmt.Sprintln(banner + "\n" + logsHelp),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logs called")
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
