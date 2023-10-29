package day02

import (
	"fmt"
	"sort"
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
	presents := parse(input)

	paper := 0
	for _, p := range presents {
		paper += p.sqft()
	}

	return fmt.Sprintf("sqft: %d", paper), nil
}

func part2(input string) (string, error) {
	presents := parse(input)

	ribbon := 0
	for _, p := range presents {
		ribbon += p.ribbon()
	}

	return fmt.Sprintf("sqft: %d", ribbon), nil
}

func parse(input string) []present {
	lines := strings.Split(input, "\n")

	var presents []present
	for _, line := range lines {
		dims := strings.Split(line, "x")

		l, _ := strconv.Atoi(dims[0])
		w, _ := strconv.Atoi(dims[1])
		h, _ := strconv.Atoi(dims[2])

		dimsInt := []int{l, w, h}
		sort.Slice(dimsInt, func(i, j int) bool {
			return dimsInt[i] < dimsInt[j]
		})

		presents = append(presents, present{l: dimsInt[0], w: dimsInt[1], h: dimsInt[2]})
	}

	return presents
}

type present struct {
	l int
	w int
	h int
}

func (p present) sqft() int {
	s1 := p.l * p.w
	smallest := s1

	s2 := p.w * p.h
	s3 := p.h * p.l

	return 2*s1 + 2*s2 + 2*s3 + smallest
}

func (p present) ribbon() int {
	o := p.l*2 + p.w*2
	e := p.l * p.w * p.h

	return o + e
}
