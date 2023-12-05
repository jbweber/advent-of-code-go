package day05

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
	seeds, h := parse(input)

	var locs []int
	for _, seed := range seeds {
		locs = append(locs, h.ToLocation(seed))
	}

	lowest := -1
	for _, loc := range locs {
		if lowest < 0 || loc < lowest {
			lowest = loc
		}
	}

	return fmt.Sprint(lowest), nil
}

func part2(input string) (string, error) {
	return "", nil
}

func parse(input string) ([]string, holder2) {
	lines := strings.Split(input, "\n")

	seeds := strings.Fields(strings.Split(lines[0], ":")[1])

	nextData := 3
	var seedToSoil []mapping

	for i := nextData; i < len(lines); i++ {
		if lines[i] == "" {
			nextData = i + 2
			break
		}

		data := strings.Fields(lines[i])
		a, _ := strconv.Atoi(data[0])
		b, _ := strconv.Atoi(data[1])
		c, _ := strconv.Atoi(data[2])

		seedToSoil = append(seedToSoil, mapping{
			DestinationRangeStart: a,
			SourceRangeStart:      b,
			RangeLength:           c,
		})

		//for j := 0; j < c; j++ {
		//	seedToSoil2[strconv.Itoa(b+j)] = strconv.Itoa(a + j)
		//}
	}

	var soilToFertilizer []mapping

	for i := nextData; i < len(lines); i++ {
		if lines[i] == "" {
			nextData = i + 2
			break
		}

		data := strings.Fields(lines[i])
		a, _ := strconv.Atoi(data[0])
		b, _ := strconv.Atoi(data[1])
		c, _ := strconv.Atoi(data[2])

		soilToFertilizer = append(soilToFertilizer, mapping{
			DestinationRangeStart: a,
			SourceRangeStart:      b,
			RangeLength:           c,
		})

		//for j := 0; j < c; j++ {
		//	soilToFertilizer2[strconv.Itoa(b+j)] = strconv.Itoa(a + j)
		//}
	}

	var fertilizerToWater []mapping

	for i := nextData; i < len(lines); i++ {
		if lines[i] == "" {
			nextData = i + 2
			break
		}

		data := strings.Fields(lines[i])
		a, _ := strconv.Atoi(data[0])
		b, _ := strconv.Atoi(data[1])
		c, _ := strconv.Atoi(data[2])

		fertilizerToWater = append(fertilizerToWater, mapping{
			DestinationRangeStart: a,
			SourceRangeStart:      b,
			RangeLength:           c,
		})

	}

	var waterToLight []mapping

	for i := nextData; i < len(lines); i++ {
		if lines[i] == "" {
			nextData = i + 2
			break
		}

		data := strings.Fields(lines[i])
		a, _ := strconv.Atoi(data[0])
		b, _ := strconv.Atoi(data[1])
		c, _ := strconv.Atoi(data[2])

		waterToLight = append(waterToLight, mapping{
			DestinationRangeStart: a,
			SourceRangeStart:      b,
			RangeLength:           c,
		})

	}

	var lightToTemp []mapping

	for i := nextData; i < len(lines); i++ {
		if lines[i] == "" {
			nextData = i + 2
			break
		}

		data := strings.Fields(lines[i])
		a, _ := strconv.Atoi(data[0])
		b, _ := strconv.Atoi(data[1])
		c, _ := strconv.Atoi(data[2])

		lightToTemp = append(lightToTemp, mapping{
			DestinationRangeStart: a,
			SourceRangeStart:      b,
			RangeLength:           c,
		})

	}

	var tempToHumidity []mapping

	for i := nextData; i < len(lines); i++ {
		if lines[i] == "" {
			nextData = i + 2
			break
		}

		data := strings.Fields(lines[i])
		a, _ := strconv.Atoi(data[0])
		b, _ := strconv.Atoi(data[1])
		c, _ := strconv.Atoi(data[2])

		tempToHumidity = append(tempToHumidity, mapping{
			DestinationRangeStart: a,
			SourceRangeStart:      b,
			RangeLength:           c,
		})

	}

	var humidityToLocation []mapping

	for i := nextData; i < len(lines); i++ {
		if lines[i] == "" {
			nextData = i + 2
			break
		}

		data := strings.Fields(lines[i])
		a, _ := strconv.Atoi(data[0])
		b, _ := strconv.Atoi(data[1])
		c, _ := strconv.Atoi(data[2])

		humidityToLocation = append(humidityToLocation, mapping{
			DestinationRangeStart: a,
			SourceRangeStart:      b,
			RangeLength:           c,
		})

	}

	return seeds, holder2{seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation}
}

type mapping struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

// in = 57
// 57 >= 50 && 57 <= 98
// 98-57 =

func (m mapping) checkMap(in int) (int, bool) {
	if in >= m.SourceRangeStart && in <= m.SourceRangeStart+m.RangeLength {
		mappedTo := (in - m.SourceRangeStart) + m.DestinationRangeStart
		return mappedTo, true
	}

	return -1, false
}

type holder2 struct {
	seedToSoil   []mapping
	soilToFert   []mapping
	fertToWater  []mapping
	waterToLight []mapping
	lightToTemp  []mapping
	tempToHumid  []mapping
	humidToLoc   []mapping
}

func (h holder2) ToLocation(in string) int {
	sn, _ := strconv.Atoi(in)

	soil := sn
	for _, m := range h.seedToSoil {
		r, ok := m.checkMap(sn)
		if ok {
			soil = r
			break
		}
	}

	fert := soil
	for _, m := range h.soilToFert {
		r, ok := m.checkMap(soil)
		if ok {
			fert = r
			break
		}
	}

	water := fert
	for _, m := range h.fertToWater {
		r, ok := m.checkMap(fert)
		if ok {
			water = r
			break
		}
	}

	light := water
	for _, m := range h.waterToLight {
		r, ok := m.checkMap(water)
		if ok {
			light = r
			break
		}
	}

	temp := light
	for _, m := range h.lightToTemp {
		r, ok := m.checkMap(light)
		if ok {
			temp = r
			break
		}
	}

	humid := temp
	for _, m := range h.tempToHumid {
		r, ok := m.checkMap(temp)
		if ok {
			humid = r
			break
		}
	}

	loc := humid
	for _, m := range h.humidToLoc {
		r, ok := m.checkMap(humid)
		if ok {
			loc = r
			break
		}
	}

	return loc
}
