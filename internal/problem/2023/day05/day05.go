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

func parse(input string) ([]string, holder) {
	lines := strings.Split(input, "\n")

	seeds := strings.Fields(strings.Split(lines[0], ":")[1])

	nextData := 3
	var seedToSoil []mapping
	seedToSoil2 := map[string]string{}
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

		for j := 0; j < c; j++ {
			seedToSoil2[strconv.Itoa(b+j)] = strconv.Itoa(a + j)
		}
	}

	var soilToFertilizer []mapping
	soilToFertilizer2 := map[string]string{}
	for i := nextData; i < len(lines); i++ {
		if lines[i] == "" {
			nextData = i + 2
			break
		}

		fmt.Println(lines[i])
		data := strings.Fields(lines[i])
		a, _ := strconv.Atoi(data[0])
		b, _ := strconv.Atoi(data[1])
		c, _ := strconv.Atoi(data[2])

		soilToFertilizer = append(soilToFertilizer, mapping{
			DestinationRangeStart: a,
			SourceRangeStart:      b,
			RangeLength:           c,
		})

		for j := 0; j < c; j++ {
			soilToFertilizer2[strconv.Itoa(b+j)] = strconv.Itoa(a + j)
		}
	}

	var fertilizerToWater []mapping
	fertilizerToWater2 := map[string]string{}
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

		for j := 0; j < c; j++ {
			fertilizerToWater2[strconv.Itoa(b+j)] = strconv.Itoa(a + j)
		}
	}

	var waterToLight []mapping
	waterToLight2 := map[string]string{}
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

		for j := 0; j < c; j++ {
			waterToLight2[strconv.Itoa(b+j)] = strconv.Itoa(a + j)
		}
	}

	var lightToTemp []mapping
	lightToTemp2 := map[string]string{}
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

		for j := 0; j < c; j++ {
			lightToTemp2[strconv.Itoa(b+j)] = strconv.Itoa(a + j)
		}
	}

	var tempToHumidity []mapping
	tempToHumidity2 := map[string]string{}
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
		for j := 0; j < c; j++ {
			tempToHumidity2[strconv.Itoa(b+j)] = strconv.Itoa(a + j)
		}
	}

	var humidityToLocation []mapping
	humidityToLocation2 := map[string]string{}
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
		for j := 0; j < c; j++ {
			humidityToLocation2[strconv.Itoa(b+j)] = strconv.Itoa(a + j)
		}
	}

	return seeds, holder{seedToSoil2, soilToFertilizer2, fertilizerToWater2, waterToLight2, lightToTemp2, tempToHumidity2, humidityToLocation2}
}

type mapping struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

type holder struct {
	seedToSoil   map[string]string
	soilToFert   map[string]string
	fertToWater  map[string]string
	waterToLight map[string]string
	lightToTemp  map[string]string
	tempToHumid  map[string]string
	humidToLoc   map[string]string
}

func (h holder) ToLocation(in string) int {
	soil, ok := h.seedToSoil[in]
	if !ok {
		soil = in
	}

	fert, ok := h.soilToFert[soil]
	if !ok {
		fert = soil
	}

	water, ok := h.fertToWater[fert]
	if !ok {
		water = fert
	}

	light, ok := h.waterToLight[water]
	if !ok {
		light = water
	}

	temp, ok := h.lightToTemp[light]
	if !ok {
		temp = light
	}

	humid, ok := h.tempToHumid[temp]
	if !ok {
		humid = temp
	}

	loc, ok := h.humidToLoc[humid]
	if !ok {
		loc = humid
	}

	l, _ := strconv.Atoi(loc)

	return l
}
