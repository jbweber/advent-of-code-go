package day05

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jbweber/advent-of-code-go/internal"
)

func Execute(input string) (string, string, error) {
	result1, err := part1(input)
	if err != nil {
		return internal.ReturnError(err)
	}

	result2, err := part2(input)
	if err != nil {
		return internal.ReturnError(err)
	}

	return result1, result2, nil
}

func part1(input string) (string, error) {
	vents, err := parseInput(input)
	if err != nil {
		return "", err
	}

	xMax := 0
	yMax := 0

	for _, vent := range vents {
		if vent.start.x > xMax {
			xMax = vent.start.x
		}

		if vent.end.x > xMax {
			xMax = vent.end.x
		}

		if vent.start.y > yMax {
			yMax = vent.start.y
		}

		if vent.end.y > yMax {
			yMax = vent.end.y
		}
	}

	grid := getGrid(xMax, yMax)

	for _, vent := range vents {
		if vent.start.x == vent.end.x {
			start := vent.start.y
			end := vent.end.y

			delta := end - start
			if delta < 0 {
				start = start + end
				end = start - end
				start = start - end
			}

			for i := start; i <= end; i++ {
				grid[i][vent.start.x] += 1
			}
		}

		if vent.start.y == vent.end.y {
			start := vent.start.x
			end := vent.end.x

			delta := end - start
			if delta < 0 {
				start = start + end
				end = start - end
				start = start - end
			}

			for i := start; i <= end; i++ {
				grid[vent.start.y][i] += 1
			}
		}

	}

	count := 0

	for i := 0; i <= yMax; i++ {
		for j := 0; j <= xMax; j++ {
			if grid[i][j] >= 2 {
				count += 1
			}
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func getGrid(xMax, yMax int) [][]int {
	grid := make([][]int, yMax+1)
	for idx := range grid {
		grid[idx] = make([]int, xMax+1)
	}

	return grid
}

func part2(input string) (string, error) {
	vents, err := parseInput(input)
	if err != nil {
		return "", err
	}

	xMax := 0
	yMax := 0

	for _, vent := range vents {
		if vent.start.x > xMax {
			xMax = vent.start.x
		}

		if vent.end.x > xMax {
			xMax = vent.end.x
		}

		if vent.start.y > yMax {
			yMax = vent.start.y
		}

		if vent.end.y > yMax {
			yMax = vent.end.y
		}
	}

	grid := getGrid(xMax, yMax)

	for _, vent := range vents {
		if vent.start.x == vent.end.x {
			start := vent.start.y
			end := vent.end.y

			delta := end - start
			if delta < 0 {
				start = start + end
				end = start - end
				start = start - end
			}

			for i := start; i <= end; i++ {
				grid[i][vent.start.x] += 1
			}

			continue
		}

		if vent.start.y == vent.end.y {
			start := vent.start.x
			end := vent.end.x

			delta := end - start
			if delta < 0 {
				start = start + end
				end = start - end
				start = start - end
			}

			for i := start; i <= end; i++ {
				grid[vent.start.y][i] += 1
			}

			continue
		}

		startX := vent.start.x
		endX := vent.end.x

		deltaX := endX - startX
		if deltaX < 0 {
			deltaX = -1
		} else {
			deltaX = 1
		}

		startY := vent.start.y
		endY := vent.end.y

		deltaY := endY - startY
		if deltaY < 0 {
			deltaY = -1
		} else {
			deltaY = 1
		}

		fmt.Printf("%d,%d -> %d,%d\n", startX, startY, endX, endY)

		count := 0
		for i := startY; i <= endY; i++ {
			fmt.Printf("%d,%d\n", i, startX+count)
			grid[i][startX+count] += 1
			count += 1
		}
	}

	count := 0

	for i := 0; i <= yMax; i++ {
		for j := 0; j <= xMax; j++ {
			if grid[i][j] >= 2 {
				count += 1
			}
		}
	}

	for idx := range grid {
		fmt.Println(grid[idx])
	}
	fmt.Println()

	return fmt.Sprintf("%d", count), nil
}

type point struct {
	x int
	y int
}

func newPoint(input string) (point, error) {
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return point{}, errors.New("didn't find two parts in input")
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return point{}, err
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return point{}, err
	}

	return point{x: x, y: y}, nil
}

type ventLine struct {
	start point
	end   point
}

func (v ventLine) String() string {
	return fmt.Sprintf("VentLine: %d,%d -> %d,%d", v.start.x, v.start.y, v.end.x, v.end.y)
}

func parseInput(input string) ([]ventLine, error) {
	lines := strings.Split(input, "\n")

	var vents []ventLine

	for _, line := range lines {

		se := strings.Split(line, " -> ")

		start, err := newPoint(se[0])
		if err != nil {
			return nil, err
		}

		end, err := newPoint(se[1])
		if err != nil {
			return nil, err
		}
		vents = append(vents, ventLine{start: start, end: end})
	}

	return vents, nil
}
