package day8

import (
	"fmt"
	"strings"

	"github.com/vincentharrius/aoc21/utils"
)

type Entry struct {
	Input  []string
	Output []string
}

func SevenSegmentSearch() int {
	lines, _ := utils.ReadFileLineString("./day8/data8.txt")

	entries := []Entry{}

	for _, line := range lines {
		l := strings.Split(line, " | ")
		input := strings.Split(l[0], " ")
		output := strings.Split(l[1], " ")

		input = input[:len(input)-1]

		entries = append(entries, Entry{Input: input, Output: output})
	}

	ones := 0
	fours := 0
	sevens := 0
	eights := 0

	for _, e := range entries {
		for _, output := range e.Output {
			fmt.Println(len(output))
			if len(output) == 2 {
				ones++
			}

			if len(output) == 3 {
				sevens++
			}

			if len(output) == 4 {
				fours++
			}

			if len(output) == 7 {
				eights++
			}
		}
	}

	fmt.Println("Ones:", ones, "\tFours:", fours, "\tSevens:", sevens, "\tEights:", eights)

	return ones + fours + sevens + eights
}
