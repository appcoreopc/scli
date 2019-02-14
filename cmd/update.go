package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")

		ExecuteUpdateCmd()
	},
}

func init() {

	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ExecuteUpdateCmd() {

	fmt.Println("Execute update cmd")

	// settingModel := &Model.CliSettingsModel{}
	// settingModel.CommandPath = "command.cli" // viper.Get("CommandPath").(string)
	// settingModel.CurrentVersion = 1.0
	// settingModel.InstallVersion = 1.0
	// settingModel.ServiceUrl = "http://test.com/command.cli.zip"
	// svc := Services.CliService{}
	// svc.Execute(settingModel)
}
