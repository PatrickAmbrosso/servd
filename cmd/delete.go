package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const deleteHelp = "Remove a service from being managed by servd"

var deleteConfirm bool

var deleteCmd = &cobra.Command{
	Use:   "delete <service-name>",
	Short: deleteHelp,
	Long:  fmt.Sprintln(banner + "\n" + deleteHelp),
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	f := deleteCmd.Flags()
	f.BoolVarP(&deleteConfirm, "confirm", "c", false, "Force delete the service")
	rootCmd.AddCommand(deleteCmd)
}
