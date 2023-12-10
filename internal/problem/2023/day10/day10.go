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

	n1, _, _ := findStartPoints2(grid, s)

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
	grid := strings.Split(input, "\n")

	s, err := findStartLocation(grid)
	if err != nil {
		return "", err
	}

	n1, _, startShape := findStartPoints2(grid, s)

	points := findShape(s, n1, grid)

	line := grid[s.Y]
	line = strings.Replace(line, "S", startShape, 1)
	grid[s.Y] = line

	count := countWraps(grid, points)

	return fmt.Sprint(count), nil
}

func findShape(start, next Point, grid []string) []Point {
	points := []Point{start, next}

	from := start
	current := next
	count := 1
	for {
		next := findNext(from, current, grid)
		from = current
		current = next
		count += 1
		if next == start {
			break
		}
		points = append(points, next)
	}

	return points
}

func drawGrid(grid []string) {
	for _, line := range grid {
		fmt.Println(line)
	}
}

func countWraps(grid []string, points []Point) int {
	count := 0
	for y := len(grid) - 1; y >= 0; y-- {
		crossed := 0
		for x := len(grid[y]) - 1; x >= 0; x-- {
			r := grid[y][x]
			// we're using the idea of "winding number" and "ray casting"
			// https://en.wikipedia.org/wiki/Point_in_polygon
			if containsX(points, Point{X: x, Y: y}) {
				if r == '|' || r == 'L' || r == 'J' {
					crossed += 1
				}
				continue
			}

			if isOdd(crossed) {
				count += 1
			}

		}
	}

	return count
}

func drawGrid2(grid []string, points []Point) {
	for y, line := range grid {
		for x, r := range line {
			if containsX(points, Point{x, y}) {
				fmt.Print("X")
			} else {
				fmt.Print(string(r))
			}
		}
		fmt.Println()
	}
}

func drawGrid3(grid []string, points []Point, wrapMap map[string]int) {
	for y, line := range grid {
		for x, _ := range line {
			if containsX(points, Point{x, y}) {
				fmt.Print("X")
			} else {
				v, ok := wrapMap[fmt.Sprintf("%d-%d", x, y)]
				if !ok {
					panic("uhoh")
				}
				if isOdd(v) {
					//fmt.Print("O")
					fmt.Print(v)
				} else {
					fmt.Print(".")
				}

			}
		}
		fmt.Println()
	}
}

func drawGrid4(grid []string, mm map[string]int) {
	for y, line := range grid {
		for x, _ := range line {
			v, ok := mm[fmt.Sprintf("%d-%d", x, y)]
			if !ok {
				fmt.Print(string(grid[y][x]))
			} else {
				fmt.Print(v)
			}
		}
		fmt.Println()
	}
}

func containsX[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func isOdd(n int) bool {
	return n%2 == 1
}
