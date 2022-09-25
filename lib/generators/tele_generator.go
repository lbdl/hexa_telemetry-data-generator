package generators

import (
	"fmt"
	"github.com/lbdl/hexa_tele/lib/types"
	"math"
)

type timeLine struct {
}

func GenerateTelemetry(tl types.TimeLine) {
	for k, v := range tl.TimeLines {
		fmt.Println("Generating for: ", k)
		genData(v)
		idxs := genOffsets(v)
		for _, v := range idxs {
			fmt.Printf("Offset index: %v\n", math.Abs(v))
		}
	}
}

func genData(tl types.TLine) {
	fmt.Println("Val: ", tl)
}

func genOffsets(tl types.TLine) []float64 {
	var offsets []float64

	for _, v := range tl.Events {
		offset := float64(tl.MaxIntervals) * (v.StartOffset / 100)
		offsets = append(offsets, offset)
	}
	return offsets
}
