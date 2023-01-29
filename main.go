package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var jsonInput string

var rootCmd = &cobra.Command{
	Use:   "jsonformatter",
	Short: "A tool to format json files",
	Long:  `A command line tool to format json files`,
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := ioutil.ReadFile(jsonInput)
		var jsonData interface{}
		json.Unmarshal([]byte(file), &jsonData)
		formattedData, _ := json.MarshalIndent(jsonData, "", "    ")
		fmt.Println(string(formattedData))
	},
}

func init() {
	rootCmd.Flags().StringVarP(&jsonInput, "file", "f", "", "json file to format")
	rootCmd.MarkFlagRequired("file")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
