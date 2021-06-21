package binary_watch

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) []string {
	result := make([]string, 0)
	for h := uint8(0); h < uint8(12); h++ {
		for m := uint8(0); m < uint8(60); m++ {
			if bits.OnesCount8(h)+bits.OnesCount8(m) == turnedOn {
				result = append(result, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return result
}
