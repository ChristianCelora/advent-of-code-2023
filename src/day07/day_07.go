package main

import (
	"adventcode/reader"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var cards_values = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	// 'J': 11, // step 1
	'J': 0, // step 2
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

type PokerHand struct {
	hand      string
	hand_type int
	bid       int
}

func CamelGameSort(poker_hands []PokerHand) func(int, int) bool {
	return func(i, j int) bool {

		if poker_hands[i].hand_type == poker_hands[j].hand_type {
			for k := 0; k < len(poker_hands[i].hand); k++ {
				v1 := GetCardValue(poker_hands[i].hand[k])
				v2 := GetCardValue(poker_hands[j].hand[k])
				if v1 != v2 {
					return v1 < v2
				}
			}
			return true
		}
		return poker_hands[i].hand_type < poker_hands[j].hand_type
	}
}

func GetCardValue(card byte) int {
	val, ok := cards_values[card]
	if !ok {
		panic("card not recognized")
	}
	return val
}

func GetPokerHandFromLine(line string) PokerHand {
	hand := strings.Split(line, " ")
	if len(hand) != 2 {
		panic("Error while parsing poker hand")
	}
	bid, err := strconv.Atoi(hand[1])
	if err != nil {
		panic(err)
	}

	return PokerHand{
		hand:      hand[0],
		hand_type: GetHandType(hand[0]),
		bid:       bid,
	}
}

type CardFreq struct {
	card      byte
	frequency int
}

/**
 * Hand Types:
 * - High Card: 0
 * - One pair: 1
 * - Two pair: 2
 * - Three of a kind: 3
 * - Full house: 4
 * - Four of a kind: 5
 * - Five of a kind: 6
 */
func GetHandType(hand string) int {
	var max_same_cards, second_max_same_cards int
	var n_jokers int

	count_pairs := make(map[byte]int)
	for i := 0; i < len(hand); i++ {
		// step 1
		// count_pairs[hand[i]]++

		// step 2
		if hand[i] == 'J' {
			n_jokers++
		} else {
			count_pairs[hand[i]]++
		}
	}

	var cards []CardFreq
	for c, f := range count_pairs {
		cards = append(cards, CardFreq{card: c, frequency: f})
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].frequency < cards[j].frequency
	})

	if len(cards) > 0 {
		max_same_cards = cards[len(cards)-1].frequency
	}
	if len(cards) > 1 {
		second_max_same_cards = cards[len(cards)-2].frequency
	}

	max_same_cards = max_same_cards + n_jokers // step 2

	switch max_same_cards {
	case 5:
		return 6
	case 4:
		return 5
	case 3:
		if second_max_same_cards > 1 {
			return 4
		} else {
			return 3
		}
	case 2:
		if second_max_same_cards > 1 {
			return 2
		} else {
			return 1
		}
	default:
		return 0
	}
}

func main() {
	var sum_winning int
	var hands []PokerHand
	lines := reader.ReadLines("./day07/data/input1_2.txt")
	for i := 0; i < len(lines); i++ {
		hands = append(hands, GetPokerHandFromLine(lines[i]))
	}

	sort.Slice(hands, CamelGameSort(hands))

	// fmt.Printf("%v\n", hands)
	for idx, hand := range hands {
		fmt.Printf("%s\n", hand.hand)
		// fmt.Printf("%d, %s - %d\n", idx, hand.hand, hand.bid)
		sum_winning += hand.bid * (idx + 1)
	}
	fmt.Printf("Sum of winnings is %d\n", sum_winning)
}
