package day16

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

var things = map[string]int{"children": 3, "cats": 7, "samoyeds": 2, "pomeranians": 3, "akitas": 0, "vizslas": 0, "goldfish": 5, "trees": 3, "cars": 2, "perfumes": 1}

func part1(input string) (string, error) {
	sueMap := parseInput(input)

	var keys []string
	for k, _ := range sueMap {
		keys = append(keys, k)
	}

	for _, key := range keys {
		sue, ok := sueMap[key]
		if !ok {
			continue
		}

		for k, v := range things {
			st, ok := sue[k]
			if !ok {
				continue
			}
			if st != v {
				delete(sueMap, key)
				break
			}
		}

	}

	fmt.Println(sueMap)

	return "", nil
}

func part2(input string) (string, error) {
	sueMap := parseInput(input)

	var keys []string
	for k, _ := range sueMap {
		keys = append(keys, k)
	}

	for _, key := range keys {
		sue, ok := sueMap[key]
		if !ok {
			continue
		}

		for k, v := range things {
			st, ok := sue[k]
			if !ok {
				continue
			}

			if k == "cats" || k == "trees" {
				if st <= v {
					delete(sueMap, key)
					break
				}
				continue
			}

			if k == "pomeranians" || k == "goldfish" {
				if st >= v {
					delete(sueMap, key)
					break
				}
				continue
			}

			if st != v {
				delete(sueMap, key)
				break
			}
		}

	}

	fmt.Println(sueMap)

	return "", nil
}

func parseInput(in string) map[string]map[string]int {
	re := regexp.MustCompile(`^Sue (\d+): ([a-z]+): (\d+), ([a-z]+): (\d+), ([a-z]+): (\d+)$`)

	sueMap := map[string]map[string]int{}

	for _, line := range strings.Split(in, "\n") {
		matches := re.FindStringSubmatch(line)
		sue := matches[1]

		x1 := matches[2]
		y1, _ := strconv.Atoi(matches[3])
		x2 := matches[4]
		y2, _ := strconv.Atoi(matches[5])
		x3 := matches[6]
		y3, _ := strconv.Atoi(matches[7])

		sueMap[sue] = map[string]int{x1: y1, x2: y2, x3: y3}
	}

	return sueMap
}
