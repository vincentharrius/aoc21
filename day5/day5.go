package day5

import (
	"math"
	"strconv"
	"strings"

	"github.com/vincentharrius/aoc21/utils"
)

type Point struct {
	X int
	Y int
}

type Vector struct {
	Start Point
	End   Point
}

type Location struct {
	Point
	Count int
}

func HydrothermalVenture() int {
	lines, _ := utils.ReadFileLineString("./day5/data5.txt")

	vectors := []Vector{}

	for _, l := range lines {
		s := strings.Split(l, " -> ")
		vectors = append(vectors, Vector{
			Start: parsePoint(s[0]),
			End:   parsePoint(s[1]),
		})
	}

	visitedLocations := []*Location{}

	// Mark out all the locations
	for _, v := range vectors {

		if !v.IsValid() {
			continue
		}

		minX := int(math.Min(float64(v.Start.X), float64(v.End.X)))
		maxX := int(math.Max(float64(v.Start.X), float64(v.End.X)))

		if minX != maxX {
			for x := minX; x <= maxX; x++ {
				found := false
				for _, loc := range visitedLocations {
					if loc.Equal(Point{X: x, Y: v.Start.Y}) {
						found = true
						loc.Count += 1
					}
				}

				if !found {
					visitedLocations = append(visitedLocations, &Location{Point: Point{X: x, Y: v.Start.Y}, Count: 1})
				}
			}
		}

		minY := int(math.Min(float64(v.Start.Y), float64(v.End.Y)))
		maxY := int(math.Max(float64(v.Start.Y), float64(v.End.Y)))

		if minY != maxY {
			for y := minY; y <= maxY; y++ {
				found := false
				for _, loc := range visitedLocations {
					if loc.Equal(Point{Y: y, X: v.Start.X}) {
						found = true
						loc.Count += 1
					}
				}

				if !found {
					visitedLocations = append(visitedLocations, &Location{Point: Point{X: v.Start.X, Y: y}, Count: 1})
				}
			}
		}
	}

	locationsVisitedMoreThanOnce := []*Location{}
	for _, loc := range visitedLocations {
		if loc.Count >= 2 {
			locationsVisitedMoreThanOnce = append(locationsVisitedMoreThanOnce, loc)
		}
	}

	return len(locationsVisitedMoreThanOnce)
}

func parsePoint(str string) Point {
	// 1,0
	coordinates := strings.Split(str, ",")
	x, _ := strconv.ParseInt(coordinates[0], 10, 64)
	y, _ := strconv.ParseInt(coordinates[1], 10, 64)
	return Point{X: int(x), Y: int(y)}
}

func (point *Point) Equal(otherPoint Point) bool {
	return point.X == otherPoint.X && point.Y == otherPoint.Y
}

func (v *Vector) IsValid() bool {
	return v.Start.X == v.End.X || v.Start.Y == v.End.Y
}
