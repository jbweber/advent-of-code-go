package day06

import (
	"fmt"
	"regexp"
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

	commands := parse(input)
	var lights [1000][1000]int

	for _, command := range commands {
		switch command.op {
		case "turn on":
			for y := command.y1; y <= command.y2; y++ {
				for x := command.x1; x <= command.x2; x++ {
					lights[y][x] = 1
				}
			}
			break
		case "turn off":
			for y := command.y1; y <= command.y2; y++ {
				for x := command.x1; x <= command.x2; x++ {
					lights[y][x] = 0
				}
			}
			break
		case "toggle":
			for y := command.y1; y <= command.y2; y++ {
				for x := command.x1; x <= command.x2; x++ {
					cur := lights[y][x]
					if cur == 1 {
						lights[y][x] = 0
					} else {
						lights[y][x] = 1
					}
				}
			}
			break
		default:
		}
	}

	count := 0
	for _, row := range lights {
		for _, light := range row {
			if light == 1 {
				count += 1
			}
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func part2(input string) (string, error) {

	commands := parse(input)
	var lights [1000][1000]int

	for _, command := range commands {
		switch command.op {
		case "turn on":
			for y := command.y1; y <= command.y2; y++ {
				for x := command.x1; x <= command.x2; x++ {
					lights[y][x] += 1
				}
			}
			break
		case "turn off":
			for y := command.y1; y <= command.y2; y++ {
				for x := command.x1; x <= command.x2; x++ {
					lights[y][x] -= 1
					if lights[y][x] < 0 {
						lights[y][x] = 0
					}
				}
			}
			break
		case "toggle":
			for y := command.y1; y <= command.y2; y++ {
				for x := command.x1; x <= command.x2; x++ {
					lights[y][x] += 2
				}
			}
			break
		default:
		}
	}

	count := 0
	for _, row := range lights {
		for _, light := range row {
			count += light
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func parse(input string) []command {
	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`^(turn on|toggle|turn off) (\d+,\d+) through (\d+,\d+)$`)

	var results []command
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		op := matches[1]
		point1Parts := strings.Split(matches[2], ",")
		point2Parts := strings.Split(matches[3], ",")

		x1, _ := strconv.Atoi(point1Parts[0])
		y1, _ := strconv.Atoi(point1Parts[1])
		x2, _ := strconv.Atoi(point2Parts[0])
		y2, _ := strconv.Atoi(point2Parts[1])

		results = append(results, command{op: op, x1: x1, y1: y1, x2: x2, y2: y2})
	}

	return results
}

type command struct {
	op string
	x1 int
	y1 int
	x2 int
	y2 int
}
