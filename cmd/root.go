package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const banner = `
▄▀▀ ██▀ █▀▄ █ █ █▀▄ │ Lightweight, Cross-Platform
▄█▀ █▄▄ █▀▄ ▀▄▀ █▄▀ │ Service Management
`

var rootCmd = &cobra.Command{
	Use:   "servd",
	Short: "servd is a lightweight service manager",
	Long:  banner,
	Run: func(cmd *cobra.Command, args []string) {

		if version, _ := cmd.Flags().GetBool("version"); version {
			fmt.Println("servd version 0.0.1")
			return
		}

		_ = cmd.Help()
	},
}

func init() {
	f := rootCmd.Flags()
	f.BoolP("version", "v", false, "prints the version of servd")
}

// Execute runs the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
