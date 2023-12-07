package day07

import (
	"fmt"
	"log"
	"sort"
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
	lines := strings.Split(input, "\n")
	var cards []Hand
	for _, line := range lines {
		items := strings.Fields(line)
		cards = append(cards, Hand{
			Cards: items[0],
			Bid:   items[1],
		})
	}

	sort.Sort(ByCards(cards))

	total := 0
	for i, x := range cards {
		bid, _ := strconv.Atoi(x.Bid)
		total += (i + 1) * bid
	}

	return fmt.Sprint(total), nil
}

func part2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	var cards []Hand
	for _, line := range lines {
		items := strings.Fields(line)
		cards = append(cards, Hand{
			Cards: items[0],
			Bid:   items[1],
			Joker: true,
		})
	}

	sort.Sort(ByCards(cards))

	total := 0
	for i, x := range cards {
		bid, _ := strconv.Atoi(x.Bid)
		total += (i + 1) * bid
	}

	return fmt.Sprint(total), nil
}

var cardToWeight = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'1': 1,
}

var cardToWeightJoker = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 11,
	'9': 10,
	'8': 9,
	'7': 8,
	'6': 7,
	'5': 6,
	'4': 5,
	'3': 4,
	'2': 3,
	'1': 2,
	'J': 1,
}

func countRunes(in string) map[rune]int {
	counts := map[rune]int{}

	for _, r := range in {
		_, ok := counts[r]
		if !ok {
			counts[r] = 0
		}

		counts[r] += 1
	}

	return counts
}

type Hand struct {
	Cards string
	Bid   string
	Joker bool
}

type ByCards []Hand

func (a ByCards) Len() int { return len(a) }
func (a ByCards) Less(i, j int) bool {
	iHandVal := scoreHand(a[i].Cards, a[i].Joker)
	jHandVal := scoreHand(a[j].Cards, a[j].Joker)

	// if the hands aren't equal
	if iHandVal != jHandVal {
		return iHandVal < jHandVal
	}

	// if the hands are equal check the cards
	for k := 0; k < 5; k++ {
		iCVal := cardToWeight[rune(a[i].Cards[k])]
		jCVal := cardToWeight[rune(a[j].Cards[k])]

		if a[i].Joker && a[j].Joker {
			iCVal = cardToWeightJoker[rune(a[i].Cards[k])]
			jCVal = cardToWeightJoker[rune(a[j].Cards[k])]
		}

		if iCVal == jCVal {
			continue
		}

		return iCVal < jCVal
	}

	// not less
	return false
}

func (a ByCards) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func scoreHand(in string, joker bool) int {
	counts := countRunes(in)

	cl := len(counts)

	switch cl {
	// five of a kind
	// no joker rules
	case 1:
		return FiveOfAKind
	// four of a kind 4,1 || full house 2,3
	case 2:
		if joker {
			_, ok := counts['J']
			if ok {
				return FiveOfAKind
			}
		}

		for _, v := range counts {
			if v == 4 || v == 1 {
				return FourOfAKind
			}
		}

		return FullHouse
	// three of a kind || 2 pair
	case 3:
		three := false
		for _, v := range counts {
			if v == 3 {
				three = true
			}
		}

		if three {
			if joker {
				_, ok := counts['J']
				if ok {
					return FourOfAKind
				}
			}
			return ThreeOfAKind
		}

		if joker {
			jc, ok := counts['J']
			if ok {
				if jc == 2 {
					return FourOfAKind
				}
				return FullHouse
			}
		}

		return TwoPair
	// one pair
	case 4:
		if joker {
			_, ok := counts['J']
			if ok {
				return ThreeOfAKind
			}
		}
		return OnePair
	//  high card
	case 5:
		if joker {
			_, ok := counts['J']
			if ok {
				return OnePair
			}
		}
		return HighCard
	default:
		log.Fatalf("unknown count %d", cl)
	}
	return -1
}

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)
