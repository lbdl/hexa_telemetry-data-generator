package types

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

type TimeLine struct {
	TimeLines map[string]TLine `mapstructure:"timeLines"`
}

type TLine struct {
	DataFieldName   string
	DataFieldType   string
	StartTime       time.Time
	TimeInterval    time.Duration
	MaxIntervals    int
	InitialValue    float64
	DriftFactor     float64
	MaxAllowedDrift float64
	Events          map[string]Event `mapstructure:"eventBlocks"`
}

type Event struct {
	EventName       string
	StartOffset     float64
	EndOffset       float64
	DriftFactor     float64
	MaxAllowedDrift float64
	TimeInterval    time.Duration
	ResetAfterEvent bool
}

func (t TimeLine) ParseToStruct(v *viper.Viper) TimeLine {
	var C TimeLine
	err := v.Unmarshal(&C)
	cobra.CheckErr(err)
	return C
}
