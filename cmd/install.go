package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Command       string
	DisplayName   string
	Description   string
	AppDirectory  string
	RestartPolicy string
)

const installHelp = "Register and install a new service"

var installCmd = &cobra.Command{
	Use:   "install <service-name>",
	Short: installHelp,
	Long:  fmt.Sprintln(banner + "\n" + installHelp),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install called")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	f := installCmd.Flags()

	f.StringVarP(&Command, "command", "c", "", "Full command to run the service (Required)")
	f.StringVarP(&DisplayName, "display-name", "d", "", "Display name of the service")
	f.StringVarP(&Description, "description", "e", "", "Description of the service")
	f.StringVarP(&AppDirectory, "app-directory", "a", "", "Working directory of the service")
	f.StringVarP(&RestartPolicy, "restart-policy", "r", "SERVICE_AUTO_START", "Restart policy of the service")
}
