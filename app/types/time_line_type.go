package types

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

type TimeLine struct {
	TimeLine struct {
		StartTime       time.Time     `yaml:"startTime"`
		MaxIntervals    int           `yaml:"maxIntervals"`
		TimeInterval    time.Duration `yaml:"timeInterval"`
		DriftFactor     float64       `yaml:"driftFactor"`
		InitialValue    float64       `yaml:"initialValue"`
		MaxAllowedDrift float64       `yaml:"maxAllowedDrift"`
		EventBlocks     []struct {
			EventName       string        `yaml:"eventName"`
			StartOffset     float64       `yaml:"startOffset"`
			EndOffset       float64       `yaml:"endOffset"`
			DriftFactor     float64       `yaml:"driftFactor"`
			MaxAllowedDrift float64       `yaml:"maxAllowedDrift"`
			TimeInterval    time.Duration `yaml:"timeInterval"`
			ResetAfterEvent bool          `yaml:"resetAfterEvent"`
		} `yaml:"eventBlocks"`
	} `yaml:"timeLine"`
}

func (t TimeLine) ParseToStruct(v *viper.Viper) TimeLine {
	var C TimeLine
	err := v.Unmarshal(&C)
	cobra.CheckErr(err)
	return C
}
