package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const stopHelp = "Stop a running service"

var stopCmd = &cobra.Command{
	Use:   "stop <service-name>",
	Short: stopHelp,
	Long:  fmt.Sprintln(banner + "\n" + stopHelp),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stop called")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
