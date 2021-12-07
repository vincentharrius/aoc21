package day3

import (
	"fmt"
	"strconv"

	"github.com/vincentharrius/aoc21/utils"
)

func CalculatePowerUsage() int {
	lines := readlines()

	binaryMap := []int{}

	for i := range lines[0] {
		binaryMap = append(binaryMap, i-i)
	}

	for _, line := range lines {
		for i, v := range line {
			if string(v) == "1" {
				binaryMap[i] += 1
			}
		}
	}

	numLines := len(lines)
	finalString := ""

	for _, v := range binaryMap {
		if v < numLines/2 {
			finalString = finalString + "0"
		} else {
			finalString = finalString + "1"
		}
	}
	invertedFinalString := ""

	for _, v := range finalString {
		if string(v) == "1" {
			invertedFinalString += "0"
		} else {
			invertedFinalString += "1"
		}
	}

	epsilon, _ := strconv.ParseInt(finalString, 2, 64)
	gamma, _ := strconv.ParseInt(invertedFinalString, 2, 64)

	return int(epsilon * gamma)
}

func VeryifyLifeSupportRating() int {

	lines := readlines()

	modes := []struct {
		Ascending bool
		Rating    int64
	}{
		{false, 0},
		{true, 0},
	}

	for mi, mode := range modes {
		matchingString := ""

		matchingLines := lines
		for i := 0; i < len(lines[0]); i++ {
			newMatchingLines := []string{}
			ones := 0
			zeroes := 0
			for _, l := range matchingLines {
				// if matchString doesnt match the line continue
				if len(matchingString) > 0 && matchingString != l[:len(matchingString)] {
					continue
				}

				// this line is a match
				newMatchingLines = append(newMatchingLines, l)

				if l[i] == '1' {
					ones += 1
				} else {
					zeroes += 1
				}
			}

			if ones < zeroes == mode.Ascending {
				matchingString += "0"
			} else {
				matchingString += "1"
			}

			if len(newMatchingLines) > 0 {
				matchingLines = newMatchingLines
			}
		}
		fmt.Println(matchingLines)
		modes[mi].Rating, _ = strconv.ParseInt(matchingLines[0], 2, 64)
	}

	fmt.Println(modes)
	return int(modes[0].Rating * modes[1].Rating)
}

func readlines() []string {
	lines, err := utils.ReadFileLineString("./day3/data3.txt")

	if err != nil {
		fmt.Println(err.Error())
	}

	return lines
}
