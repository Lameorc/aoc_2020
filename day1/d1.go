package day1

import (
	"log"
	"strconv"
	"strings"
)

// Solve the day
func Solve(lines []string) int {
	out := 0
	nLines := len(lines)

	asInts := make([]int32, nLines)
	for i, line := range lines {
		r, err := strconv.ParseInt(strings.TrimSpace(line), 0, 32)
		if err != nil {
			log.Fatalf("Failed to parse %s: %s", line, err)
		}

		asInts[i] = int32(r)
	}

	// part 1
	// for i := 0; i < nLines; i++ {
	// 	for j := i + 1; j < nLines; j++ {
	// 		if asInts[i]+asInts[j] == 2020 {
	// 			log.Printf("Numbers of interest should be %d and %d", asInts[i], asInts[j])
	// 			out = int(asInts[i] * asInts[j])
	// 			break
	// 		}
	// 		if out != 0 {
	// 			break
	// 		}
	// 	}
	// }

	// part 2
	for i := 0; i < nLines; i++ {
		for j := i + 1; j < nLines; j++ {
			for k := j + 1; k < nLines; k++ {
				if asInts[i]+asInts[j]+asInts[k] == 2020 {
					log.Printf("Numbers of interest should be %d, %d and %d", asInts[i], asInts[j], asInts[k])
					out = int(asInts[i] * asInts[j] * asInts[k])
					break
				}

			}
			if out != 0 {
				break
			}
		}
		if out != 0 {
			break
		}
	}

	return out

}
