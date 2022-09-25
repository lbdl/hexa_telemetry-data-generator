package generators

import (
	"fmt"
	"github.com/lbdl/hexa_tele/lib/types"
)

func GenerateTelemetry(tele types.TimeLine) {
	for k, v := range tele.TimeLines {
		fmt.Println("Generating for: ", k)
		genData(v)
	}
}

func genData(tl types.TLine) {
	fmt.Println("Val: ", tl)
}
