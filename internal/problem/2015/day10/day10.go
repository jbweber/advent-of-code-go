package day10

import (
	"fmt"
	"os"
	"runtime/pprof"
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
	f, err := os.Create("cpu.profile")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()
	digs := parse(input)

	result := make([]int, len(digs))
	copy(result, digs)

	for i := 0; i < 40; i++ {
		result = do1(result)
		fmt.Printf("after %d runs ", i+1)
	}

	return fmt.Sprintf("%v", len(result)), nil
}

func do1(in []int) []int {
	if len(in) == 0 {
		return nil
	}

	var result []int

	// single digit case
	if len(in) == 1 {
		result = append(result, 1, in[0])
		return result
	}

	// two different
	first := in[0]
	second := in[1]

	if first != second {
		result = append(result, 1, first)

		result = append(result, do1(in[1:])...)
		return result
	}

	// same total of two
	if len(in) == 2 {
		result = append(result, 2, first)
		return result
	}

	third := in[2]

	if first != third {
		result = append(result, 2, first)
		result = append(result, do1(in[2:])...)
		return result
	}

	// three digits
	result = append(result, 3, first)
	result = append(result, do1(in[3:])...)

	//
	return result
}

func do2(in string) string {
	//if len(in) == 0 {
	//	return ""
	//}
	//
	//if len(in) == 1 {
	//	return fmt.Sprintf("1%s", in)
	//}

	curr := in[0]
	currCount := 1
	//tail := ""
	tail := strings.Builder{}

	for i := 1; i < len(in); i++ {
		if in[i] == curr {
			currCount += 1
		} else {
			//tail = fmt.Sprintf("%s%d%c", tail, currCount, curr)
			//tail += strconv.Itoa(currCount) + string(curr)
			tail.WriteString(strconv.Itoa(currCount))
			tail.WriteString(string(curr))
			curr = in[i]
			currCount = 1
		}
	}

	tail.WriteString(strconv.Itoa(currCount))
	tail.WriteString(string(curr))

	return tail.String()
}

func part2(input string) (string, error) {
	//f, err := os.Create("cpu.profile")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//
	//if err := pprof.StartCPUProfile(f); err != nil {
	//	panic(err)
	//}
	//defer pprof.StopCPUProfile()

	for i := 0; i < 50; i++ {
		input = do2(input)
	}

	return fmt.Sprintf("%v", len(input)), nil
}

func parse(input string) []int {
	input = strings.Trim(input, "\n")

	var digs []int
	for _, i := range input {
		// digits are 48 (0) to 57 (9) in ascii char so subtracting zero gives real numeral
		digs = append(digs, int(i-'0'))
	}

	return digs
}
