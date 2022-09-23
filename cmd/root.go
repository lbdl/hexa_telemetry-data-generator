/*
Copyright © 2022 lbdl: timstorey@hexaponics.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	dbCfgFile string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "hx-gen",
		Short: "Reads from a given config file or terminal flags and then creates a set of time based test data",
		Long: `Given a config file will read the values and then generate test data

	This can then be passed into a data base to create a set of test data.`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Global flags and configuration settings.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (defaults to [$HOME/hexa-gen.yml, $CWD/configs/hexa-gen.yml)")
	rootCmd.PersistentFlags().StringVarP(&dbCfgFile, "dbconf", "d", "", "db config file (defaults to [$HOME/db_conf.yaml, $CWD/configs/db_conf.yml)")
}

func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
