package generators

import (
	"fmt"
	"github.com/lbdl/hexa_tele/lib/types"
	"golang.org/x/exp/slices"
	"math"
	"time"
)

type eventOffset struct {
	startOffset int
	endOffset   int
}

// this is a datapoint and can contain either
// a single eventAtom or an array.
// this way we can handle the offset time duration
// variation that we may see inside an event block
type timeAtom struct {
	offset int
	atoms  []eventAtom
}

type eventAtom struct {
	unitDuration time.Duration
	unitValue    float64
}

func GenerateTelemetry(tl types.TimeLine) {
	for k, v := range tl.TimeLines {
		fmt.Println("Generating for: ", k)
		genData(v)
		idxs := genEventOffsets(v)
		for _, v := range idxs {
			fmt.Printf("Offset index: %v\n", v)
		}
	}
}

func genData(tl types.TLine) {
	fmt.Println("Val: ", tl)
}

func genEventOffsets(tl types.TLine) []eventOffset {
	var offsets []eventOffset
	for _, v := range tl.Events {
		offSet := eventOffset{}
		offSet.startOffset = int(math.Abs((float64(tl.DataPoints) * (v.StartOffset / 100))))
		offSet.endOffset = int(math.Abs(float64(tl.DataPoints) * (v.EndOffset / 100)))
		offsets = append(offsets, offSet)
	}
	return offsets
}

// check that we don't have an overlap of offsets
// if we do then just set the start to the last + 1
// and note this in the log/stdout/stderr.
func validateEventStartOffsets(e []eventOffset) {
	for i, ev := range e {
		if i > 0 {
			if ev.startOffset <= e[i].endOffset {
				ev.startOffset = e[i].endOffset + 1
			}
		}
	}
}

// if the end index is smaller than the start index just
// delete the struct from the slice and note this in the log
// stderr/stdout.
func validateEventEndOffsets(e []eventOffset) {
	for i, ev := range e {
		if ev.startOffset <= ev.endOffset {
			e = slices.Delete(e, i, i+1)
		}
	}
}
