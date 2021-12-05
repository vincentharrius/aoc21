package day2

import "github.com/vincentharrius/aoc21/utils"

type Direction string

const (
	Up      Direction = "up"
	Down    Direction = "down"
	Forward Direction = "forward"
)

type Movement struct {
	Direction Direction `json:"direction"`
	Length    int       `json:"length"`
}

func Dive() int {
	horizontal := 0
	depth := 0
	aim := 0

	movements := []Movement{}

	utils.ReadJson("./day2/data2.json", &movements)

	for _, m := range movements {
		if m.Direction == Up {
			aim -= m.Length
		} else if m.Direction == Down {
			aim += m.Length
		} else if m.Direction == Forward {
			horizontal += m.Length
			depth += aim * m.Length
		}
	}

	return horizontal * depth
}
