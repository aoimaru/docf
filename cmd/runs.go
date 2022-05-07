/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/aoimaru/docf/lib"
	"github.com/spf13/cobra"
)

type Optionss struct {
	esc  bool
	file bool
}

var (
	opts = &Optionss{}
)

type OutPuts struct {
	Datas []lib.OutPut `json:datas`
}

// runsCmd represents the runs command
var runsCmd = &cobra.Command{
	Use:   "runs",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		outputs := make([]lib.OutPut, 0)
		// fmt.Println("runs called")
		if len(args) == 0 {
			log.Fatal("specify root Directory")
		}
		rootDic := args[0]
		fps, err := lib.ListFiles(rootDic)
		if err != nil {
			log.Fatal(err)
		}
		if opts.file {
			lib.CreateDirectory(fps)
			for _, fp := range fps {
				if strings.HasSuffix(fp, "Dockerfile") {
					fmt.Println(fp, opt.esc)
					output := lib.GetRun(fp, opt.esc)
					fmt.Println(output)
					fmt.Println(fp)
					dirname, basename := filepath.Split(fp)
					if basename == "Dockerfile" {
						file_name := time.Now().String() + ".json"
						// fmt.Println(dirname + file_name)
						lib.CreateFile(dirname, file_name, output)
					} else {
						file_name := strings.Replace(basename, ".Dockerfile", ".json", -1)
						// fmt.Println(file_name)
						lib.CreateFile(dirname, file_name, output)
					}

				}
			}
		} else {
			for _, fp := range fps {
				if strings.HasSuffix(fp, "Dockerfile") {
					fmt.Println(fp, opt.esc)
					output := lib.GetRun(fp, opt.esc)
					fmt.Println(reflect.TypeOf(output))
					outputs = append(outputs, output)
				}
			}
			results := OutPuts{outputs}
			fmt.Println(results)
			jsonData, err := json.Marshal(results)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("%s\n", jsonData)
		}

	},
}

func init() {
	rootCmd.AddCommand(runsCmd)
	runsCmd.Flags().BoolVarP(&opts.esc, "esc", "e", false, "escape sequence")
	runsCmd.Flags().BoolVarP(&opts.file, "file", "f", false, "escape sequence")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
