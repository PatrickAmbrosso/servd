package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const infoHelp = "Display system and platform details for debugging"

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: infoHelp,
	Long:  fmt.Sprintln(banner + "\n" + infoHelp),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info called")
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
