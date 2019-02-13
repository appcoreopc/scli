// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/appcoreopc/scli/Model"

	"github.com/appcoreopc/scli/Services"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Verbose bool

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:              "root",
	TraverseChildren: true,
	Short:            "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called")
	},
}

func init() {

	fmt.Println("executing main root cmd ")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.MarkFlagRequired("verbose")

	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//viper.AddRemoteProvider("etcd", "http://127.0.0.1:2389", "/v2/keys/message")
	//viper.SetConfigType("json") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop"
	//viper.SetConfigName("settings")
	//viper.ReadRemoteConfig()
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if Verbose == true {
		fmt.Println("setting to verbose mode")
	}

	fmt.Println("value", viper.Get("name"))

	settingModel := &Model.CliSettingsModel{}
	settingModel.CommandPath = "command.cli" // viper.Get("CommandPath").(string)
	settingModel.CurrentVersion = 1.0
	settingModel.InstallVersion = 1.0
	settingModel.ServiceUrl = "http://test.com/command.cli.zip"
	svc := Services.CliService{}

	svc.Execute(settingModel)

}
