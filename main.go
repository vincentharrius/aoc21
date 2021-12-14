package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vincentharrius/aoc21/day1"
	"github.com/vincentharrius/aoc21/day2"
	"github.com/vincentharrius/aoc21/day3"
	"github.com/vincentharrius/aoc21/day4"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("No args")
		return
	}
	skipnext := false
	for i, arg := range os.Args {
		if skipnext {
			skipnext = false
			continue
		}

		switch arg {
		case "-c":
			createDayFolder(os.Args[i+1])
			skipnext = true
		case "1":
			fmt.Println(day1.IncreasedMeasurements(1))
			fmt.Println(day1.IncreasedMeasurements(3))
			break

		case "2":
			fmt.Println(day2.Dive())
			break

		case "3":
			fmt.Println(day3.CalculatePowerUsage())
			fmt.Println(day3.VeryifyLifeSupportRating())
			break

		case "4":
			fmt.Println(day4.ThatsABingo())
			break
		}
	}
}

func createDayFolder(day string) {

	foldername := "day" + day
	if err := os.Mkdir(foldername, 0755); err != nil {
		log.Fatal(err)
	}

	if _, err := os.Create(foldername + "/day" + day + ".go"); err != nil {
		log.Fatal(err)
	}
}
