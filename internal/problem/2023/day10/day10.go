package day10

import (
	"fmt"
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
	grid := strings.Split(input, "\n")

	x, y := findStartLocation(grid)
	fmt.Println(x, y)

	x1, y1, x2, y2 := findStartPoints(grid, x, y)

	s := Point{X: x, Y: y}
	n1 := Point{X: x1, Y: y1}
	_ = Point{X: x2, Y: y2}

	from := s
	current := n1
	count := 1
	for {
		next := findNext(from, current, grid)
		from = current
		current = next
		count += 1
		if next == s {
			break
		}
	}

	return fmt.Sprint(count / 2), nil
}

func findStartLocation(in []string) (int, int) {
	y := 0
	x := 0

	for i, line := range in {
		if strings.Contains(line, "S") {
			x = strings.Index(line, "S")
			y = i
			return x, y
		}
	}

	return -1, -1
}

type Point struct {
	X int
	Y int
}

func findNext(from Point, current Point, grid []string) Point {
	pu := Point{X: current.X, Y: current.Y - 1}
	pd := Point{X: current.X, Y: current.Y + 1}
	pl := Point{X: current.X - 1, Y: current.Y}
	pr := Point{X: current.X + 1, Y: current.Y}

	//fmt.Println(string(right), string(left), string(down), string(up))

	if grid[current.Y][current.X] == '|' {
		if from == pu {
			return pd
		}

		return pu
	}

	if grid[current.Y][current.X] == '-' {
		if from == pl {
			return pr
		}

		return pl
	}

	if grid[current.Y][current.X] == 'L' {
		if from == pu {
			return pr
		}

		return pu
	}

	if grid[current.Y][current.X] == 'J' {
		if from == pu {
			return pl
		}

		return pu
	}

	if grid[current.Y][current.X] == '7' {
		if from == pd {
			return pl
		}

		return pd
	}

	if grid[current.Y][current.X] == 'F' {
		if from == pd {
			return pr
		}

		return pd
	}

	panic("cannot get here findNext")

	return Point{}
}

func findStartPoints(grid []string, x, y int) (int, int, int, int) {
	lenY := len(grid)
	lenX := len(grid[0])

	right, left, down, up := '.', '.', '.', '.'
	// could be x+1, y
	if x < lenX-1 {
		right = rune(grid[y][x+1])
	}
	// could be x-1, y
	if x > 0 {
		left = rune(grid[y][x-1])
	}
	// could be x, y+1
	if y < lenY-1 {
		down = rune(grid[y+1][x])
	}
	// could be x, y-1
	if y > 0 {
		up = rune(grid[y-1][x])
	}

	// |
	if contains([]rune{'|', 'F', '7'}, up) && contains([]rune{'|', 'L', 'J'}, down) {
		return x, y + 1, x, y - 1
	}

	// -
	if contains([]rune{'-', 'F', 'L'}, left) && contains([]rune{'-', 'J', '7'}, right) {
		return x - 1, y, x + 1, y
	}

	// L
	if contains([]rune{'|', 'F', '7'}, up) && contains([]rune{'-', 'J', '7'}, right) {
		return x, y + 1, x + 1, y
	}

	// J
	if contains([]rune{'|', 'F', '7'}, up) && contains([]rune{'-', 'L', 'F'}, left) {
		return x, y + 1, x - 1, y
	}

	// 7
	if contains([]rune{'-', 'F', 'L'}, left) && contains([]rune{'|', 'L', 'J'}, down) {
		return x - 1, y, x, y - 1
	}

	// F
	if contains([]rune{'-', '7', 'J'}, right) && contains([]rune{'|', 'L', 'J'}, down) {
		return x + 1, y, x, y - 1
	}

	fmt.Println(string(right), string(left), string(down), string(up))

	panic("cannot get here findStart")
}

func contains(in []rune, v rune) bool {
	for _, s := range in {
		if v == s {
			return true
		}
	}
	return false
}

func part2(input string) (string, error) {
	return "", nil
}
