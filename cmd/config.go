package cmd

import (
	"fmt"

	"github.com/PatrickAmbrosso/servd/database"
	"github.com/spf13/cobra"
)

const configHelp = "Configure servd settings & preferences"

var (
	listConfig           bool
	resetConfig          bool
	confirmReset         bool
	configExportPath     string
	configLogLevel       string
	configServiceTimeout int
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: configHelp,
	Long:  banner + "\n" + configHelp,
	Run: func(cmd *cobra.Command, args []string) {
		if listConfig && cmd.Flags().Changed("list") {
			fmt.Println("Listing configuration values...")
			return
		}

		if resetConfig && cmd.Flags().Changed("reset") {
			fmt.Println("Resetting configuration values... | confirm:", confirmReset)
		}

		config := database.Config{} // make it a get method to get existing config

		if configExportPath != "" && cmd.Flags().Changed("configExportPath") {
			fmt.Println("Exporting configuration servd configurations")
		}

		if configLogLevel != "" && cmd.Flags().Changed("configLogLevel") {
			config.LogLevel = configLogLevel
		}

		if configServiceTimeout != 0 && cmd.Flags().Changed("configServiceTimeout") {
			config.ServiceTimeout = configServiceTimeout
		}
	},
}

func init() {
	f := configCmd.Flags()

	f.BoolVarP(&listConfig, "list", "l", false, "Show current configuration values")
	f.BoolVarP(&resetConfig, "reset", "r", false, "Reset configuration to defaults")
	f.BoolVarP(&confirmReset, "confirm", "c", false, "Confirm configuration reset (used with --reset)")
	f.StringVarP(&configExportPath, "export", "e", "config.yaml", "Set export path for configuration")
	f.StringVarP(&configLogLevel, "log-level", "L", "info", "Set log level (debug, info, warn, error)")
	f.IntVarP(&configServiceTimeout, "service-timeout", "t", 30, "Set default service timeout (in seconds)")

	rootCmd.AddCommand(configCmd)

}
