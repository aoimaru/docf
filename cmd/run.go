/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aoimaru/docf/lib"
	"github.com/spf13/cobra"
)

type Options struct {
	esc    bool
	origin bool
}

var (
	opt = &Options{}
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("specify Dockerfile")
		}
		output := lib.GetRun(args[0], opt.esc)
		// fmt.Println(output)
		jsonData, err := json.Marshal(output)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s\n", jsonData)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolVarP(&opt.esc, "esc", "e", false, "escape sequence")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
