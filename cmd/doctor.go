package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const doctorHelp = "Diagnose and fix common issues in the environment"

var diagnosticsExportPath string

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: doctorHelp,
	Long:  fmt.Sprintln(banner + "\n" + doctorHelp),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("doctor called")
	},
}

func init() {
	f := doctorCmd.Flags()
	f.StringVarP(&diagnosticsExportPath, "export", "e", "", "Export path for the diagnostics file to be saved")
	rootCmd.AddCommand(doctorCmd)
}
