/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"encoding/json"

	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"github.com/sunnyvale-it/helm-graph/chart"
)

var Name string
var Version string
var Repo string
var OutputFormat string

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		chart := &chart.Chart{}

		chart.Name = Name
		chart.Version = Version
		chart.Repo = Repo

		chart.Graph()

		if OutputFormat == "json" {
			json, err := json.Marshal(chart)
			if err != nil {
				fmt.Printf("An error occured: %v\n", err)
				return
			}
			fmt.Println(string(json))
		}

		if OutputFormat == "yaml" {
			yaml, err := yaml.Marshal(chart)
			if err != nil {
				fmt.Printf("An error occured: %v\n", err)
				return
			}
			fmt.Println(string(yaml))
		}

	},
}

func init() {
	rootCmd.AddCommand(renderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	renderCmd.Flags().StringVarP(&Name, "name", "n", "", "The chart name (required)")
	renderCmd.MarkFlagRequired("name")

	renderCmd.Flags().StringVarP(&Version, "version", "v", "", "The chart version (required)")
	renderCmd.MarkFlagRequired("version")

	renderCmd.Flags().StringVarP(&Repo, "repo", "r", "", "The chart repo in URI format (required)")
	renderCmd.MarkFlagRequired("repo")

	renderCmd.Flags().StringVarP(&OutputFormat, "output", "o", "json", "The graph output format (json / yaml)")

}
