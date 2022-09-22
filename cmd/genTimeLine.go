/*
Copyright © 2022 lbdl: timstorey@hexaponics.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// genTimeStampsCmd represents the genTimeStamps command
var (
	v            *viper.Viper
	startTime    string
	maxIntervals int
	baseInterval string
	driftFactor  float64
	maxDrift     float64
	dType        []map[string]interface{}

	genTimeLineCmd = &cobra.Command{
		Use:   "genTimeLine",
		Short: "Generates a timeline from a config file",
		Long: `Uses a config yml file to generate a time line data
	structure.

	Useage: hx-gen genTimeLine [-c, --config] myConfig.yaml/.yml
	
	If no config file is passed will look in the default dirs as follow:
	$HOME/hx-gen-config.yml
	$PWD/hx-gen-config.yml`,
		Run: func(cmd *cobra.Command, args []string) {
			printConf()
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			v = viper.New()
			readConf(v)
			parseConf(v)
		},
	}
)

func init() {
	// add pFlags here if needed
	rootCmd.AddCommand(genTimeLineCmd)
}

func printConf() {
	fmt.Println("max Intervals: ", maxIntervals)
	fmt.Println("startTime: ", startTime)
	fmt.Println("interval: ", baseInterval)
	fmt.Println("drift: ", driftFactor)
	fmt.Println("data: ", dType)
}

func parseConf(v *viper.Viper) {
	cfgFile = v.ConfigFileUsed()
	startTime = v.GetString("timeLine.startTime")
	maxIntervals = v.GetInt("timeLine.maxIntervals")
	baseInterval = v.GetString("timeLine.timeInterval")
	driftFactor = v.GetFloat64("timeLine.driftFactor")
}

func readConf(v *viper.Viper) {
	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		pwd, err := os.UserHomeDir()
		home, err := os.Getwd()
		cobra.CheckErr(err)

		v.AddConfigPath(home)
		v.AddConfigPath(pwd)
		v.SetConfigType("yml")

		// default config which we really should
		// get from the flasg but we aren't right now
		v.SetConfigName("hexa-gen")
		v.AddConfigPath("../configs/")
	}

	if err := v.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config:", v.ConfigFileUsed())
	} else {
		fmt.Println("No config file found")
	}
}
