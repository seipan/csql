// MIT License

// Copyright (c) 2023 Yamasaki Shotaro

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/seipan/csql/lib"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "csql",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
	     ___________ ____    __ 
	    / ____/ ___// __ \  / / 
	   / /    \__ \/ / / / / /  
	  / /___ ___/ / /_/ / / /___
	  \____//____/\___\_\/_____/
									  
								   
	`)
	},
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Println(err)
		}
		dsn, err := cmd.Flags().GetString("dsn")
		if err != nil {
			log.Println(err)
		}
		types, err := cmd.Flags().GetString("type")
		if err != nil {
			log.Println(err)
		}
		query, err := cmd.Flags().GetBool("query")
		if err != nil {
			log.Println(err)
		}
		check, err := cmd.Flags().GetBool("check")
		if err != nil {
			log.Println(err)
		}
		cfg, _ := lib.ParseYML(".csql.yaml")
		cfg = replaceConfig(cfg, dsn, types, path)
		err = checkConfig(*cfg)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		if check {
			err = lib.CsvFormatExec(*cfg)
			if err == nil {
				fmt.Println("csv format is correct")
			} else {
				fmt.Println("csv format is incorrect :", err)
				os.Exit(1)
			}
		} else if query {
			str, err := lib.QueryExec(*cfg)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			fmt.Println(str)
		} else {
			err = lib.InsertExec(*cfg)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.csql.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("check", "c", false, "check csv format")
	rootCmd.Flags().BoolP("query", "q", false, "output query")
	rootCmd.Flags().StringP("path", "p", "", "FilePath for Parsing CSVFile")
	rootCmd.Flags().StringP("dsn", "d", "", "DSN for Connecting Database")
	rootCmd.Flags().StringP("type", "t", "", "Database Type")
}

func checkConfig(cfg lib.Config) error {
	if cfg.Type == "" {
		return errors.New("type is empty")
	}
	if cfg.DSN == "" {
		return errors.New("dsn is empty")
	}
	if cfg.Filepath == "" {
		return errors.New("filepath is empty")
	}
	return nil
}

func replaceConfig(cfg *lib.Config, dns string, types string, path string) *lib.Config {
	if cfg == nil {
		cfg = &lib.Config{}
	}
	cfg.DSN = replaceString(cfg.DSN, dns)
	cfg.Type = replaceString(cfg.Type, types)
	cfg.Filepath = replaceString(cfg.Filepath, path)
	return cfg
}

func replaceString(str string, target string) string {
	if str == "" {
		return target
	}
	return str
}
