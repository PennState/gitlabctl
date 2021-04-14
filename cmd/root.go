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
	"strings"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var dryRun bool

const InitialPadding = 2

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitlabctl",
	Short: "A gitlab CLI tool",
	Long:  `A tool used by software engineering to manage gitlab.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		PrintLogo()
	},
	SilenceUsage: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	log.SetHandler(cli.Default)
	cli.Default.Padding = InitialPadding
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gitlabctl.yaml)")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "log actions, don't perform gitlab api calls (for testing)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".glctl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gitlabctl")
	}

	viper.AutomaticEnv() // read in environment variables that match
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("gitlabctl")

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Config file not found, using defaults: %s\n", err)
	} else {
		fmt.Fprintf(os.Stderr, "Using config file: %s\n", viper.ConfigFileUsed())
	}
}

func PrintLogo() {
	color.Blue("          _ __  __      __         __  __")
	color.Blue("   ____ _(_) /_/ /___ _/ /_  _____/ /_/ /")
	color.Blue("  / __ `/ / __/ / __ `/ __ \\/ ___/ __/ / ")
	color.Blue(" / /_/ / / /_/ / /_/ / /_/ / /__/ /_/ /  ")
	color.Blue(" \\__, /_/\\__/_/\\__,_/_.___/\\___/\\__/_/   ")
	color.Blue("/____/                                   ")
}
