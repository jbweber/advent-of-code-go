package day07

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
	inmap := parse(input)
	wiremap := map[string]uint16{}

	for k, _ := range inmap {
		findValue(k, inmap, wiremap)
	}

	return fmt.Sprintf("a -> %d", wiremap["a"]), nil
}

func part2(input string) (string, error) {
	inmap := parse(input)
	wiremap := map[string]uint16{}

	wiremap["b"] = 16076

	for k, _ := range inmap {
		findValue(k, inmap, wiremap)
	}

	return fmt.Sprintf("a -> %d", wiremap["a"]), nil
}

func parse(input string) map[string]string {
	lines := strings.Split(input, "\n")

	theMap := map[string]string{}

	for _, line := range lines {
		parts := strings.Split(line, " -> ")

		wire := parts[1]
		instruction := ""

		parts2 := strings.Split(parts[0], " ")
		switch len(parts2) {
		case 1: // num -> wire or wire -> wire
			instruction = fmt.Sprintf("SET %s", parts2[0])
			break
		case 2: // NOT wire -> wire
			instruction = fmt.Sprintf("%s %s", parts2[0], parts2[1])
			break
		case 3: //
			instruction = fmt.Sprintf("%s %s %s", parts2[1], parts2[0], parts2[2])
		}

		theMap[wire] = instruction
	}

	return theMap
}

func isInteger(v string) bool {
	if _, err := strconv.Atoi(v); err == nil {
		return true
	}
	return false
}

func findValue(wire string, inmap map[string]string, wiremap map[string]uint16) {
	instruction := inmap[wire]
	inparts := strings.Split(instruction, " ")

	var arg1 uint16
	var arg2 uint16
	switch inparts[0] {
	case "SET":
		if isInteger(inparts[1]) {
			v, _ := strconv.Atoi(inparts[1])
			arg1 = uint16(v)
		} else {
			_, ok := wiremap[inparts[1]]
			if !ok {
				findValue(inparts[1], inmap, wiremap)
			}
			arg1 = wiremap[inparts[1]]
		}
		wiremap[wire] = doOp("SET", arg1, 0)
		break
	case "NOT":
		_, ok := wiremap[inparts[1]]
		if !ok {
			findValue(inparts[1], inmap, wiremap)
		}
		arg1 = wiremap[inparts[1]]
		result := doOp("NOT", arg1, 0)
		wiremap[wire] = result
		break
	case "AND":
		if isInteger(inparts[1]) {
			v, _ := strconv.Atoi(inparts[1])
			arg1 = uint16(v)
		} else {
			_, ok := wiremap[inparts[1]]
			if !ok {
				findValue(inparts[1], inmap, wiremap)
			}
			arg1 = wiremap[inparts[1]]
		}

		_, ok := wiremap[inparts[2]]
		if !ok {
			findValue(inparts[2], inmap, wiremap)
		}
		arg2 = wiremap[inparts[2]]
		result := doOp("AND", arg1, arg2)
		wiremap[wire] = result
		break
	case "OR":
		_, ok := wiremap[inparts[1]]
		if !ok {
			findValue(inparts[1], inmap, wiremap)
		}
		arg1 = wiremap[inparts[1]]
		_, ok = wiremap[inparts[2]]
		if !ok {
			findValue(inparts[2], inmap, wiremap)
		}
		arg2 = wiremap[inparts[2]]
		result := doOp("OR", arg1, arg2)
		wiremap[wire] = result
		break
	case "LSHIFT":
		_, ok := wiremap[inparts[1]]
		if !ok {
			findValue(inparts[1], inmap, wiremap)
		}
		arg1 = wiremap[inparts[1]]
		arg2i, _ := strconv.Atoi(inparts[2])
		arg2 = uint16(arg2i)
		result := doOp("LSHIFT", arg1, arg2)
		wiremap[wire] = result
		break
	case "RSHIFT":
		_, ok := wiremap[inparts[1]]
		if !ok {
			findValue(inparts[1], inmap, wiremap)
		}
		arg1 = wiremap[inparts[1]]
		arg2i, _ := strconv.Atoi(inparts[2])
		arg2 = uint16(arg2i)
		result := doOp("RSHIFT", arg1, arg2)
		wiremap[wire] = result
		break
	}
}

func doOp(operation string, x, y uint16) uint16 {
	switch operation {
	case "SET":
		return x
	case "NOT":
		return ^x
	case "AND":
		return x & y
	case "OR":
		return x | y
	case "LSHIFT":
		return x << y
	case "RSHIFT":
		return x >> y
	}
	return 0
}

// id string
// signal uint16
//
