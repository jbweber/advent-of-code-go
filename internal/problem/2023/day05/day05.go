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

	result2, err := part3(input)
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
	seeds, h := parse(input)

	minLoc := -1
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		count := seeds[i+1]
		for j := start; j < start+count; j++ {
			if minLoc == -1 {
				minLoc = h.ToLocation(j)
			} else {
				// old logic was bad because it was trying to make a HUGE slice which was slowing things down
				// much better to be not lame and just check minimum as we get to it
				minLoc = min(minLoc, h.ToLocation(j))
			}
		}
	}

	return fmt.Sprint(minLoc), nil
}

func part3(input string) (string, error) {
	seeds, mappingStack := parse2(input)

	results := RangeToLocation(seeds, mappingStack)

	minVal := results[0].Start

	for _, result := range results {
		minVal = min(minVal, result.Start)
	}

	return fmt.Sprint(minVal), nil
}

func RangeToLocation(in []internal.Range, mappingStack [][]mapping) []internal.Range {
	for _, stack := range mappingStack {
		in = doMapping(in, stack)
	}

	return in
}

//func RangeToLocation(r1 internal.Range, stack [][]mapping, queue chan internal.Range) internal.Range {
//	var results []internal.Range
//
//	for _, ms := range stack {
//		result := doMapping([]internal.Range{r1}, ms)
//		//for _, s := range ms {
//		//	r2 := internal.Range{Start: s.SourceRangeStart, End: s.SourceRangeStart + s.RangeLength}
//		//
//		//	rOut, leftovers, overlap := internal.FindPartialIntersection(r1, r2)
//		//	if !overlap {
//		//		fmt.Println(r1, r2, rOut, leftovers, overlap)
//		//		continue
//		//	}
//		//
//		//	os := (rOut.Start - s.SourceRangeStart) + s.DestinationRangeStart
//		//	oe := (rOut.End - s.SourceRangeStart) + s.DestinationRangeStart
//		//
//		//	r1 = internal.Range{Start: os, End: oe}
//		//
//		//	fmt.Println(r1, r2, rOut, leftovers, overlap)
//		//	for _, leftover := range leftovers {
//		//		queue <- leftover
//		//	}
//		//	break
//		//}
//	}
//
//	return r1
//}

func doMapping(in []internal.Range, stack []mapping) []internal.Range {
	queue := make(chan internal.Range, len(in)*10)
	for _, seed := range in {
		queue <- seed
	}

	var results []internal.Range

	for len(queue) > 0 {
		seed, ok := <-queue
		if !ok {
			fmt.Println("queue is empty")
			break
		}

		for _, s := range stack {
			r2 := internal.Range{Start: s.SourceRangeStart, End: s.SourceRangeStart + s.RangeLength}

			rOut, leftovers, overlap := internal.FindPartialIntersection(seed, r2)
			// no mapping has happened
			if !overlap {
				continue
			}

			os := (rOut.Start - s.SourceRangeStart) + s.DestinationRangeStart
			oe := (rOut.End - s.SourceRangeStart) + s.DestinationRangeStart

			seed = internal.Range{Start: os, End: oe}

			for _, leftover := range leftovers {
				queue <- leftover
			}
			break
		}

		results = append(results, seed)
	}

	return results
}

func parse2(input string) ([]internal.Range, [][]mapping) {
	lines := strings.Split(input, "\n")

	seedsFields := strings.Fields(strings.Split(lines[0], ":")[1])

	var seeds []internal.Range
	for i := 0; i < len(seedsFields); i += 2 {
		s, _ := strconv.Atoi(seedsFields[i])
		r, _ := strconv.Atoi(seedsFields[i+1])
		seeds = append(seeds, internal.Range{Start: s, End: s + r})
	}

	var mappingStack [][]mapping

	nextData := 3

	cur := 0
	for cur < len(lines) {
		var curMapping []mapping
		for i := nextData; i < len(lines); i++ {
			if lines[i] == "" {
				nextData = i + 2
				break
			}

			data := strings.Fields(lines[i])
			a, _ := strconv.Atoi(data[0])
			b, _ := strconv.Atoi(data[1])
			c, _ := strconv.Atoi(data[2])

			curMapping = append(curMapping, mapping{
				DestinationRangeStart: a,
				SourceRangeStart:      b,
				RangeLength:           c,
			})

			cur = i + 1
		}

		mappingStack = append(mappingStack, curMapping)
	}

	return seeds, mappingStack
}

func parse(input string) ([]int, holder2) {
	lines := strings.Split(input, "\n")

	seeds := strings.Fields(strings.Split(lines[0], ":")[1])

	var seedsInt []int
	for _, seed := range seeds {
		s, _ := strconv.Atoi(seed)
		seedsInt = append(seedsInt, s)
	}

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

	return seedsInt, holder2{seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation}
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

func (h holder2) ToLocation(in int) int {
	soil := in
	for _, m := range h.seedToSoil {
		r, ok := m.checkMap(in)
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
