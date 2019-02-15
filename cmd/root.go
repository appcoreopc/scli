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
	"log"
	"os"

	"github.com/appcoreopc/scli/Fops"

	"github.com/appcoreopc/scli/Model"
	"github.com/appcoreopc/scli/Providers"
	"github.com/appcoreopc/scli/Services"
	"github.com/spf13/cobra"
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

		// Get value from consul
		kv := new(Providers.ConsulClient)
		cmdJson := kv.GetKeyJson("tools").(Model.CommandCliModel)

		// Read local version info
		jsonReader := new(Fops.JsonReader)
		ivInfo := jsonReader.GetCommandJson("command.cli").(Model.CommandCliModel)

		// tenary : haha
		installedVersion := ivInfo.Version

		if args == nil {
			ShowToolsHelp(&cmdJson.Tools)
		} else if len(args) > 0 {

			for _, element := range cmdJson.Tools {

				if element.Name == args[0] {
					// check local for exe //
					log.Println("Execute tooling command")

					//
					cliService := new(Services.CliService)
					cliService.Execute(installedVersion, &element)
				}
			}
		}
	},
}

func init() {

	//rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	//rootCmd.MarkFlagRequired("verbose")

	//viper.SetConfigName("config") // name of config file (without extension)
	//viper.AddConfigPath(".")      // optionally look for config in the working directory
	//err := viper.ReadInConfig()   // Find and read the config file
	// if err != nil { // Handle errors reading the config file
	// 	panic(fmt.Errorf("Fatal error config file: %s \n", err))
	// }

	// viper.AddRemoteProvider("etcd", "http://127.0.0.1:2379", "v2/keys/settings")
	// viper.SetConfigType("json") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop"
	// viper.SetConfigName("settings")
	// viper.ReadRemoteConfig()
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ShowToolsHelp(tools *[]Model.ToolModel) {

	//jr := Fops.JsonReader{}

	//cmdJson := jr.GetCommandJson("command.cli/command.cli.json").(Model.CommandCliModel)

	fmt.Println("Tools available")
	fmt.Println("Usage: scli <toolname> <parameter1> <parameter2> ")
	fmt.Println("Example scli new --name NewProjectName")
	fmt.Println("--------------------------------------------------------")
	for _, element := range *tools {
		fmt.Println("Tool name ", element.Name)
		fmt.Println("Description ", element.Description)
		fmt.Println("Version", element.Version)
		fmt.Println("Commandline(<toolname>)", element.Command)
		fmt.Println("--------------------------------------------------------")
	}
}
