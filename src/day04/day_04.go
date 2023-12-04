package main

import (
	"adventcode/reader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Scratchcard struct {
	id              int
	winning_numbers []int
	user_numbers    []int
	instances       int
}

func SliceIndexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func GetScratchcardId(line string) int {
	line_card_id := strings.TrimSpace(strings.Replace(line, "Card", "", -1))
	card_id, err := strconv.Atoi(line_card_id)
	if err != nil {
		panic(err)
	}

	return card_id
}

func GetScratchcardFromLine(line string) Scratchcard {
	var card Scratchcard

	numbers := strings.Split(line, "|")
	if len(numbers) != 2 {
		panic("wrong split line: " + line)
	}

	card = Scratchcard{
		winning_numbers: getNumbersSliceFromString(numbers[0]),
		user_numbers:    getNumbersSliceFromString(numbers[1]),
		instances:       1,
	}

	return card
}

func getNumbersSliceFromString(line string) []int {
	var res []int
	numbers_string := strings.TrimSpace(line)
	numbers := strings.Split(numbers_string, " ")
	for _, n_string := range numbers {
		if n_string == "" {
			continue
		}
		n, err := strconv.Atoi(n_string)
		if err != nil {
			panic(err)
		}
		res = append(res, n)
	}
	return res
}

func CalcScratchCardPoints(card Scratchcard) int {
	found := countScratchcardWinningNumbers(card)
	if found == 0 {
		return 0
	}
	return int(math.Pow(2, float64(found-1)))
}

func countScratchcardWinningNumbers(card Scratchcard) int {
	var found int

	for _, user_n := range card.user_numbers {
		if SliceIndexOf(user_n, card.winning_numbers) != -1 {
			found++
		}
	}

	return found
}

const (
	// MAX_SCRATCHCARD = 6 // test
	MAX_SCRATCHCARD = 201
)

func main() {
	var sum_points int
	var sum_instances int
	var scratchcard_instances [MAX_SCRATCHCARD]Scratchcard
	lines := reader.ReadLines("./day04/data/input1_2.txt")
	for _, line := range lines {
		line_card := strings.Split(line, ":")
		if len(line_card) != 2 {
			panic("wrong split line: " + line)
		}

		scratchcard := GetScratchcardFromLine(line_card[1])
		scratchcard.id = GetScratchcardId(line_card[0]) - 1

		scratchcard_instances[scratchcard.id] = scratchcard
		sum_points += CalcScratchCardPoints(scratchcard)
	}
	// step 1
	fmt.Printf("The sum of the card points is %d\n", sum_points)

	// step 2
	for i := 0; i <= len(scratchcard_instances)-1; i++ {
		scratchcard := scratchcard_instances[i]
		instances_won := countScratchcardWinningNumbers(scratchcard)
		fmt.Printf("id: %d, won: %d, mult: %d\n", scratchcard.id, instances_won, scratchcard.instances)
		for j := 0; j <= instances_won-1; j++ {
			card_won_id := scratchcard.id + 1 + j
			scratchcard_instances[card_won_id].instances = scratchcard_instances[card_won_id].instances + (1 * scratchcard.instances)
		}
	}

	for _, scratchcard := range scratchcard_instances {
		fmt.Printf("id: %d, instances: %d\n", scratchcard.id, scratchcard.instances)

		sum_instances += scratchcard.instances
	}

	fmt.Printf("The sum of the card instances is %d\n", sum_instances)

}
