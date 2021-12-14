package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vincentharrius/aoc21/utils"
)

const RowHeight = 5
const RowWidth = 5

type BingoCard struct {
	C []*Cell
}

type Cell struct {
	V int
	M bool
}

func ThatsABingo() int {
	lines, _ := utils.ReadFileLineString("./day4/data4.txt")

	numbers := strings.Split(lines[0], ",")
	lines = lines[1:]

	bingoCards := []*BingoCard{}

	for i := 1; i < len(lines); i += (RowHeight + 1) {
		bingoCards = append(bingoCards, parseBingoCard(lines[i:i+RowHeight]))
	}

	for _, number := range numbers {
		n := toInt(number)

		for _, b := range bingoCards {
			b.Mark(n)
			if sum := b.GetBingoSum(); sum != 0 {
				fmt.Println("Sum:", sum, "n:", n)
				return sum * n
			}
		}
	}

	return 0
}

func parseBingoCard(lines []string) *BingoCard {
	bingoCard := BingoCard{}

	for _, row := range lines {
		numbers := strings.Split(row, " ")
		for _, number := range numbers {
			if number == "" {
				continue
			}

			bingoCard.C = append(bingoCard.C, &Cell{V: toInt(number), M: false})
		}
	}

	return &bingoCard
}

func toInt(number string) int {
	n, _ := strconv.ParseInt(number, 10, 64)
	return int(n)
}

func (b *BingoCard) Mark(number int) {
	for _, c := range b.C {
		if c.V == number {
			c.M = true
			return
		}
	}
}

func (b *BingoCard) GetBingoSum() int {

	// Checks horizontal bingo
	for i := 0; i < len(b.C); i += 5 {
		cells := b.C[i : i+5]

		bingo := true
		for _, c := range cells {
			bingo = bingo && c.M
		}

		if bingo {
			sum := 0
			for _, unmarked := range b.Unmarked() {
				sum += unmarked.V
			}
			return sum
		}
	}

	for i := 0; i < 5; i++ {
		bingo := true
		for j, c := range b.C {
			if j%5 != 0 {
				continue
			}

			bingo = bingo && c.M
		}

		if bingo {
			sum := 0
			for _, unmarked := range b.Unmarked() {
				sum += unmarked.V
			}
			return sum
		}
	}

	return 0
}

func (b *BingoCard) Unmarked() []*Cell {
	unmarked := []*Cell{}
	for _, c := range b.C {
		if !c.M {
			unmarked = append(unmarked, c)
		}
	}
	return unmarked
}
