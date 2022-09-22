/*
Copyright © 2022 timstorey@hexaponics.com

*/
package cmd

import (
	"fmt"
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
		Short: "Reads from a given config file or terinal flags and then creates a set of time based test data",
		Long: `Given a config file will read the values and then generate test data

	This can then be passed into a data base to create a set of test data.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
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
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (defaults to [$HOME/hexa-gen.yaml, $CWD/configs/hexa-gen.yml)")
	rootCmd.PersistentFlags().StringVarP(&dbCfgFile, "dbconf", "d", "", "db config file (defaults to [$HOME/db_conf.yaml, $CWD/configs/db_conf.yml)")
	//rootCmd.PersistentFlags().IntVarP(&maxIntervals, "maxIntervals", "I", 100, "The max number of time intervals to use")
	//viper.BindPFlag("maxIntervals", rootCmd.PersistentFlags().Lookup("timeLine.maxIntervals"))
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		pwd, err := os.Getwd()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(pwd)
		viper.SetConfigType("yml")
		viper.SetConfigName("hexa-gen")
		viper.AddConfigPath("../configs/")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("No config file found...")
	}

	// bind global vars
	cfgFile = viper.ConfigFileUsed()
	startTime = viper.GetString("timeLine.startTime")
	maxIntervals = viper.GetInt("timeLine.maxIntervals")
	baseInterval = viper.GetString("timeLine.timeInterval")
	driftFactor = viper.GetFloat64("timeLine.driftFactor")
}
