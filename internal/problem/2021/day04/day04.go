package day04

import (
	"errors"
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

func part1(input string) (string, error) {
	draws, boards, _ := parseBoards(input)

	// play first five
	for i := 0; i < 5; i++ {
		for _, board := range boards {
			board.Play(draws[i])
		}
	}

	winIdx := -1

	for idx, board := range boards {
		check := board.WinCheck()
		if check {
			winIdx = idx
			break
		}
	}

	if winIdx >= 0 {
		v := boards[winIdx].Calculate()
		return "winner " + strconv.Itoa(v), nil
	}

	for _, number := range draws[5:] {
		for idx, board := range boards {
			board.Play(number)
			if board.WinCheck() {
				v := boards[idx].Calculate()
				return "winner " + strconv.Itoa(v), nil
			}
		}
	}

	return "winner " + "unknown", nil
}

func part2(input string) (string, error) {
	draws, boards, _ := parseBoards(input)

	// play first five
	for i := 0; i < 5; i++ {
		for _, board := range boards {
			board.Play(draws[i])
		}
	}

	var won []int

	for idx, board := range boards {
		check := board.WinCheck()
		if check {
			won = append(won, idx)
		}
	}

	if len(won) == len(boards) {
		v := boards[won[len(won)-1]].Calculate()
		return "winner " + strconv.Itoa(v), nil
	}

	for _, number := range draws[5:] {
		for idx, board := range boards {
			// we won already no need to keep playing
			if contains(won, idx) {
				continue
			}

			board.Play(number)
			if board.WinCheck() {
				won = append(won, idx)
			}
		}

		if len(won) == len(boards) {
			v := boards[won[len(won)-1]].Calculate()
			return "winner " + strconv.Itoa(v), nil
		}
	}

	return "winner " + "unknown", nil
}

func contains(in []int, val int) bool {
	for _, i := range in {
		if i == val {
			return true
		}
	}
	return false
}

type bingoBoard struct {
	board      [5][5]string
	lastPlayed string
	win        bool
}

func (b *bingoBoard) Calculate() int {
	sum := 0

	for _, row := range b.board {
		for _, i := range row {
			if i != "x" {
				v, _ := strconv.Atoi(i)
				sum += v
			}
		}
	}

	return b.LastPlayed() * sum
}

func (b *bingoBoard) LastPlayed() int {
	v, _ := strconv.Atoi(b.lastPlayed)
	return v
}

func (b *bingoBoard) Play(play string) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.board[i][j] == play {
				b.board[i][j] = "x"
			}
		}
	}

	b.lastPlayed = play
}

func (b *bingoBoard) String() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("Last Played: %s\n", b.lastPlayed))
	sb.WriteString(fmt.Sprintf("Win Check: %t\n", b.win))
	for i := 0; i < 5; i++ {
		sb.WriteString("[ ")
		for j := 0; j < 5; j++ {
			item := b.board[i][j]
			if len(item) == 1 {
				sb.WriteString(" ")
				sb.WriteString(b.board[i][j])
			} else {
				sb.WriteString(b.board[i][j])
			}

			if j != 4 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString(" ]\n")
	}

	return sb.String()
}

func (b *bingoBoard) WinCheck() bool {
	for i := 0; i < 5; i++ {

		if b.board[i][0] == "x" && b.board[i][1] == "x" && b.board[i][2] == "x" && b.board[i][3] == "x" && b.board[i][4] == "x" {
			b.win = true
			return true
		}

		if b.board[0][i] == "x" && b.board[1][i] == "x" && b.board[2][i] == "x" && b.board[3][i] == "x" && b.board[4][i] == "x" {
			b.win = true
			return true
		}
	}

	b.win = false
	return false
}

func parseBoards(input string) ([]string, []*bingoBoard, error) {
	lines := strings.Split(input, "\n")
	if len(lines) < 7 {
		return nil, nil, errors.New("text input doesn't contain at least one board")
	}

	draws := strings.Split(lines[0], ",")

	boardLines := lines[2:]

	if (len(boardLines)+1)%6 != 0 {
		return nil, nil, errors.New("boards input doesn't seem to have correct contents")
	}

	var boards []*bingoBoard

	start := 0
	end := len(boardLines)

	for {
		board := [5][5]string{}

		board[0] = [5]string(spaceSplit(boardLines[start])[0:5])
		board[1] = [5]string(spaceSplit(boardLines[start+1])[0:5])
		board[2] = [5]string(spaceSplit(boardLines[start+2])[0:5])
		board[3] = [5]string(spaceSplit(boardLines[start+3])[0:5])
		board[4] = [5]string(spaceSplit(boardLines[start+4])[0:5])

		boards = append(boards, &bingoBoard{board: board})

		start += 6
		if start > end {
			break
		}
	}

	return draws, boards, nil
}

func spaceSplit(input string) []string {
	input = strings.TrimLeft(input, " ")
	input = strings.TrimRight(input, " ")

	r, _ := regexp.Compile(`\s+`)

	split := r.Split(input, -1)

	set := []string{}

	for i := range split {
		set = append(set, split[i])
	}

	return set
}
