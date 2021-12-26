package day7

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/vincentharrius/aoc21/utils"
)

func TheTreacheryOfWhales() int {
	lines, _ := utils.ReadFileLineString("./day7/data7.txt")
	numbers := []int{}

	for _, l := range strings.Split(lines[0], ",") {
		n, _ := strconv.ParseInt(l, 10, 64)
		numbers = append(numbers, int(n))
	}

	lowestFuel := 0
	lowestPos := 0

	for _, pos := range numbers {

		fuel := 0
		for _, n := range numbers {
			if n == pos {
				continue
			}

			fuel += int(math.Abs(float64(n - pos)))
		}

		if fuel < lowestFuel || lowestFuel == 0 {
			lowestFuel = fuel
			lowestPos = pos
		}
	}
	fmt.Println("Pos:", lowestPos, "\tFuel:", lowestFuel)
	return lowestFuel
}
