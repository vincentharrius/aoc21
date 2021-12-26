package day6

import (
	"strconv"
	"strings"

	"github.com/vincentharrius/aoc21/utils"
)

type Fish struct {
	N int
}

const DAYS = 80

func LanternFish() int {

	lines, _ := utils.ReadFileLineString("./day6/data6.txt")

	numbers := strings.Split(lines[0], ",")

	fishes := []*Fish{}
	// Create new fishes
	for _, number := range numbers {
		n, _ := strconv.ParseInt(number, 10, 64)
		fishes = append(fishes, &Fish{N: int(n)})
	}

	for day := 0; day <= DAYS; day++ {

		babyFishes := []*Fish{}
		for _, fish := range fishes {
			if fish.IsBirthDay() {
				babyFishes = append(babyFishes, fish.Birth())
			}

			if day == 0 {
				continue
			}

			fish.Sleep()
		}
		fishes = append(fishes, babyFishes...)
	}

	return len(fishes)
}

func (f *Fish) Sleep() {
	f.N--
}

func (f *Fish) IsBirthDay() bool {
	return f.N == 0
}

func (f *Fish) Birth() *Fish {
	f.N = 7
	return &Fish{N: 8}
}
