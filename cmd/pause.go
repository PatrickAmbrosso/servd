package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const pauseHelp = "Temporarily suspend a running service (if supported)"

var pauseCmd = &cobra.Command{
	Use:   "pause <service-name>",
	Short: pauseHelp,
	Long:  fmt.Sprintln(banner + "\n" + pauseHelp),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pause called")
	},
}

func init() {
	rootCmd.AddCommand(pauseCmd)
}
