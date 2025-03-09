package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const servicesHelp = "Manage services (list, export, import)"

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: servicesHelp,
	Long:  fmt.Sprintln(banner + "\n" + servicesHelp),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("services called")
	},
}

func init() {
	rootCmd.AddCommand(servicesCmd)
}
