/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// genTimeStampsCmd represents the genTimeStamps command
var (
	genTimeLine = &cobra.Command{
		Use:   "genTimeLine",
		Short: "Generates a timeline from a config file",
		Long: `Uses a config yml file to generate a time line data
	structure.

	Useage: hx-gen genTimeLine [-c, --config] myConfig.yaml/.yml
	
	If no config file is passed will look in the default dirs as follow:
	$HOME/hx-gen-config.yml
	$PWD/hx-gen-config.yml`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("max Intervals: ", maxIntervals)
			fmt.Println("startTime: ", startTime)
			fmt.Println("interval: ", baseInterval)
			fmt.Println("drift: ", driftFactor)
			fmt.Println("data: ", dType)
		},
	}
	startTime    string
	maxIntervals int
	baseInterval string
	driftFactor  float64
	maxDrift     float64
	dType        string
)

func init() {
	rootCmd.AddCommand(genTimeLine)

	//dType = viper.GetString("timeLine.eventDataType[0].typeName")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genTimeStampsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genTimeStampsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
