package day1

import (
	"github.com/vincentharrius/aoc21/utils"
)

func IncreasedMeasurements(pace int) int {
	depths := []int{}
	utils.ReadJson("./day1/data1.json", &depths)

	timesIncreased := 0
	previousPaceDepth := 0

	for i := 0; i <= len(depths)-pace; i++ {
		if i == 0 {
			continue
		}

		paceDepth := 0

		for j := 0; j < pace; j++ {
			paceDepth += depths[j+i]
		}

		if paceDepth > previousPaceDepth {
			timesIncreased += 1
		}

		previousPaceDepth = paceDepth
	}

	return timesIncreased
}
