package day02

import (
	"errors"
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
	s := submarine1{}

	for _, line := range strings.Split(input, "\n") {
		cmd := strings.Split(line, " ")

		val, err := strconv.Atoi(cmd[1])
		if err != nil {
			return "", err
		}

		switch cmd[0] {
		case "down":
			s.down(val)
			break
		case "forward":
			s.forward(val)
			break
		case "up":
			s.up(val)
			break
		default:
			return "", errors.New("unknown command")
		}
	}

	return strconv.Itoa(s.output()), nil
}

func part2(input string) (string, error) {
	s := submarine2{}

	for _, line := range strings.Split(input, "\n") {
		cmd := strings.Split(line, " ")

		val, err := strconv.Atoi(cmd[1])
		if err != nil {
			return "", err
		}

		switch cmd[0] {
		case "down":
			s.down(val)
			break
		case "forward":
			s.forward(val)
			break
		case "up":
			s.up(val)
			break
		default:
			return "", errors.New("unknown command")
		}
	}

	return strconv.Itoa(s.output()), nil
}

type submarine1 struct {
	x int
	y int
}

func (s *submarine1) down(v int) {
	s.y += v
}

func (s *submarine1) forward(v int) {
	s.x += v
}

func (s *submarine1) up(v int) {
	s.y -= v
}

func (s *submarine1) output() int {
	return s.x * s.y
}

type submarine2 struct {
	aim int
	x   int
	y   int
}

func (s *submarine2) down(v int) {
	s.aim += v
}

func (s *submarine2) forward(v int) {
	s.x += v
	s.y += s.aim * v
}

func (s *submarine2) up(v int) {
	s.aim -= v
}

func (s *submarine2) output() int {
	return s.x * s.y
}
