/*
Copyright © 2022 lbdl: timstorey@hexaponics.com
*/
package cmd

import (
	"fmt"
	"github.com/lbdl/hexa_tele/lib/generators"
	"github.com/lbdl/hexa_tele/lib/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// genTimeStampsCmd represents the genTimeStamps command
var (
	v  *viper.Viper
	tl types.TimeLine

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
	generators.GenerateTelemetry(tl)
}

func parseConf(v *viper.Viper) {
	cfgFile = v.ConfigFileUsed()
	tl = tl.ParseToStruct(v)
}

func readConf(v *viper.Viper) {
	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		pwd, err := os.UserHomeDir()
		home, err := os.Getwd()
		cobra.CheckErr(err)

		// some defaults and search paths
		v.AddConfigPath(home)
		v.AddConfigPath(pwd)
		v.SetConfigType("yml")
		v.SetConfigName("hexa-gen")
		v.AddConfigPath("configs/")
	}

	err := v.ReadInConfig()
	cobra.CheckErr(err)
	fmt.Fprintln(os.Stderr, "Using config:", v.ConfigFileUsed())
}
