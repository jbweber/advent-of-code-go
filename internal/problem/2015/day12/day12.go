package day12

import (
	"encoding/json"
	"fmt"
	"reflect"
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

	f := false
	b := strings.Builder{}
	for _, c := range input {
		switch c {
		case '-',
			'1',
			'2',
			'3',
			'4',
			'5',
			'6',
			'7',
			'8',
			'9',
			'0':
			if !f {
				b.WriteString(",")
			}
			b.WriteString(string(c))
			f = true
			break
		default:
			f = false
		}
	}

	total := 0
	for _, n := range strings.Split(strings.TrimLeft(b.String(), ","), ",") {
		v, _ := strconv.Atoi(n)
		total += v
	}

	return fmt.Sprintf("%d", total), nil
}

func part2(input string) (string, error) {
	var wrapper []any
	err := json.Unmarshal([]byte(input), &wrapper)
	if err != nil {
		return "", err
	}

	total := handleSlice(wrapper)

	return fmt.Sprintf("%d", total), nil
}

func handleSlice(in []any) int {
	total := 0
	for _, i := range in {
		switch v := i.(type) {
		case int:
			total += v
			break
		case float64:
			total += int(v)
			break
		case []any:
			total += handleSlice(v)
			break
		case map[string]any:
			total += handleMap(v)
			break
		case string:
			continue
		default:
			fmt.Printf("UNKNOWN: %v %v\n", reflect.TypeOf(v), v)
			break
		}
	}

	return total
}

func handleMap(in map[string]any) int {
	total := 0
	for _, v := range in {
		switch i := v.(type) {
		case int:
			total += i
			break
		case float64:
			total += int(i)
			break
		case []any:
			total += handleSlice(i)
			break
		case map[string]any:
			total += handleMap(i)
			break
		case string:
			if i == "red" {
				return 0
			}
			continue
		default:
			fmt.Printf("UNKNOWN: %v %v\n", reflect.TypeOf(v), v)
			break
		}
	}

	return total
}
