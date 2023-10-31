package day09

import (
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
	graph := parse2(input)

	var keys []string
	for k, _ := range graph {
		keys = append(keys, k)
	}

	choices := internal.Permutations(keys)

	distance := 0
	path := ""

	for idx, choice := range choices {
		option := 0
		possiblePath := ""
		for i := 0; i < len(choice)-1; i++ {
			possiblePath += fmt.Sprintf("%s -> %s\n", choice[i], choice[i+1])
			if v, ok := graph[choice[i]][choice[i+1]]; !ok {
				continue
			} else {
				option += v
			}
		}

		if idx == 0 {
			distance = option
			path = possiblePath
		}

		if option < distance {
			distance = option
			path = possiblePath
		}
	}
	fmt.Println(path)
	return fmt.Sprintf("shortest distance is %d", distance), nil
}

func part2(input string) (string, error) {
	graph := parse2(input)

	var keys []string
	for k, _ := range graph {
		keys = append(keys, k)
	}

	choices := internal.Permutations(keys)

	distance := 0
	path := ""

	for idx, choice := range choices {
		option := 0
		possiblePath := ""
		for i := 0; i < len(choice)-1; i++ {
			possiblePath += fmt.Sprintf("%s -> %s\n", choice[i], choice[i+1])
			if v, ok := graph[choice[i]][choice[i+1]]; !ok {
				continue
			} else {
				option += v
			}
		}

		if idx == 0 {
			distance = option
			path = possiblePath
		}

		if option > distance {
			distance = option
			path = possiblePath
		}
	}
	fmt.Println(path)
	return fmt.Sprintf("longest distance is %d", distance), nil
}

type edge struct {
	Vertex string
	Weight int
}

func parse(input string) map[string][]edge {
	lines := strings.Split(input, "\n")

	graph := map[string][]edge{}

	for _, line := range lines {
		parts := strings.Split(line, " = ")
		weight, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		points := strings.Split(parts[0], " to ")

		if _, ok := graph[points[0]]; !ok {
			graph[points[0]] = make([]edge, 0)
		}
		graph[points[0]] = append(graph[points[0]], edge{Vertex: points[1], Weight: weight})

		if _, ok := graph[points[1]]; !ok {
			graph[points[1]] = make([]edge, 0)
		}
		graph[points[1]] = append(graph[points[1]], edge{Vertex: points[0], Weight: weight})
	}

	return graph
}

func parse2(input string) map[string]map[string]int {
	lines := strings.Split(input, "\n")

	graph := map[string]map[string]int{}

	for _, line := range lines {
		parts := strings.Split(line, " = ")
		weight, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		points := strings.Split(parts[0], " to ")

		if _, ok := graph[points[0]]; !ok {
			graph[points[0]] = make(map[string]int, 0)
		}
		graph[points[0]][points[1]] = weight

		if _, ok := graph[points[1]]; !ok {
			graph[points[1]] = make(map[string]int, 0)
		}
		graph[points[1]][points[0]] = weight
	}

	return graph
}
