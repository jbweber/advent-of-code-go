package day14

import (
	"fmt"
	"math"
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
	deerMap := do(input)

	distance := 0
	winner := ""
	for k, v := range deerMap {
		gone := v.howFar(2503)
		if gone > distance {
			distance = gone
			winner = k
		}
	}

	return fmt.Sprintf("%s went %d", winner, distance), nil
}

func part2(input string) (string, error) {
	deerMap := do(input)

	points := map[string]int{}
	roundDistance := map[string]int{}

	for k, _ := range deerMap {
		points[k] = 0
		roundDistance[k] = 0
	}

	for i := 0; i < 2503; i++ {
		for k, v := range deerMap {
			roundDistance[k] = v.howFar(i + 1)
		}

		leaders := findLeaders(roundDistance)
		for _, l := range leaders {
			points[l] += 1
		}
	}

	mp := 0
	mpg := ""
	for k, v := range points {
		if v > mp {
			mp = v
			mpg = k
		}
	}

	return fmt.Sprintf("%s went %d", mpg, mp), nil
}

func findLeaders(in map[string]int) []string {
	lead := 0
	for _, v := range in {
		if v > lead {
			lead = v
		}
	}

	var res []string

	for k, v := range in {
		if v == lead {
			res = append(res, k)
		}
	}

	return res
}

func do(in string) map[string]deer {
	deerMap := map[string]deer{}
	for _, line := range strings.Split(in, "\n") {
		data := strings.Split(line, " ")
		name := data[0]
		speed, _ := strconv.Atoi(data[3])
		speedFor, _ := strconv.Atoi(data[6])
		restFor, _ := strconv.Atoi(data[13])

		deerMap[name] = deer{name: name, speed: speed, speedFor: speedFor, restFor: restFor}
	}

	return deerMap
}

type deer struct {
	name     string
	speed    int
	speedFor int
	restFor  int
}

func (d deer) howFar(elapsed int) int {
	section := d.speedFor + d.restFor

	if elapsed < d.speedFor {
		return elapsed * d.speed
	}

	if elapsed < section {
		return d.speedFor * d.speed
	}

	sections := int(math.Floor(float64(elapsed) / float64(section)))

	distance := sections * d.speedFor * d.speed

	left := elapsed - (sections * section)

	if left < d.speedFor {
		distance += left * d.speed
	} else {
		distance += d.speedFor * d.speed
	}

	return distance
}
