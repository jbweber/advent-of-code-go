package day10

import (
	"errors"
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

	s, err := findStartLocation(grid)
	if err != nil {
		return "", err
	}

	fmt.Println(s)

	n1, _, startShape := findStartPoints2(grid, s)

	fmt.Println(startShape)

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

func findStartLocation(in []string) (Point, error) {
	for i, line := range in {
		if strings.Contains(line, "S") {
			x := strings.Index(line, "S")
			y := i
			return Point{X: x, Y: y}, nil
		}
	}

	return Point{}, errors.New("cannot find `S` shape in starting grid")
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

func findStartPoints2(grid []string, start Point) (Point, Point, string) {
	lenY := len(grid)
	lenX := len(grid[0])

	right, left, down, up := '.', '.', '.', '.'
	// could be x+1, y
	if start.X < lenX-1 {
		right = rune(grid[start.Y][start.X+1])
	}
	// could be x-1, y
	if start.X > 0 {
		left = rune(grid[start.Y][start.X-1])
	}
	// could be x, y+1
	if start.Y < lenY-1 {
		down = rune(grid[start.Y+1][start.X])
	}
	// could be x, y-1
	if start.Y > 0 {
		up = rune(grid[start.Y-1][start.X])
	}

	// |
	if contains([]rune{'|', 'F', '7'}, up) && contains([]rune{'|', 'L', 'J'}, down) {
		return Point{start.X, start.Y + 1}, Point{start.X, start.Y - 1}, "|"
	}

	// -
	if contains([]rune{'-', 'F', 'L'}, left) && contains([]rune{'-', 'J', '7'}, right) {
		return Point{start.X - 1, start.Y}, Point{start.X + 1, start.Y}, "-"
	}

	// L
	if contains([]rune{'|', 'F', '7'}, up) && contains([]rune{'-', 'J', '7'}, right) {
		return Point{start.X, start.Y + 1}, Point{start.X + 1, start.Y}, "L"
	}

	// J
	if contains([]rune{'|', 'F', '7'}, up) && contains([]rune{'-', 'L', 'F'}, left) {
		return Point{start.X, start.Y + 1}, Point{start.X - 1, start.Y}, "J"
	}

	// 7
	if contains([]rune{'-', 'F', 'L'}, left) && contains([]rune{'|', 'L', 'J'}, down) {
		return Point{start.X - 1, start.Y}, Point{start.X, start.Y - 1}, "7"
	}

	// F
	if contains([]rune{'-', '7', 'J'}, right) && contains([]rune{'|', 'L', 'J'}, down) {
		return Point{start.X + 1, start.Y}, Point{start.X, start.Y - 1}, "F"
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
