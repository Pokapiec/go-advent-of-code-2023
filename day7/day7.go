package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type PlayerHand struct {
	Hand   string
	Weight int
}

func GetHandScore(hand string) int {
	score := 0
	charCounts := map[rune]int{}

	for _, char := range hand {
		charCounts[char]++
	}

	for _, count := range charCounts {
		score += count * count
	}

	return score
}

func ReplaceJokersWithBetter(hand string) string {
	if !strings.Contains(hand, "J") {
		return hand
	}

	charCounts := map[rune]int{}

	for _, char := range hand {
		charCounts[char]++
	}

	if strings.Contains(hand, "J") {
		maxKey := rune(0)
		maxValue := 0
		for key, count := range charCounts {
			if count > maxValue && key != 'J' {
				maxKey = key
				maxValue = count
			}
		}

		hand = strings.ReplaceAll(hand, "J", string(maxKey))
	}

	return hand
}

func SignToNumber(sign string) int {
	signMapping := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"T": 10,
		"J": 1,
	}

	if unicode.IsDigit(rune(sign[0])) {
		result, _ := strconv.Atoi(sign)
		return result
	} else {
		return signMapping[sign]
	}
}

func CompareHands(hand1 string, hand2 string) int {
	score1 := GetHandScore(hand1)
	score2 := GetHandScore(hand2)

	if score1 > score2 {
		return 1
	}

	if score1 < score2 {
		return -1
	}

	for i := 0; i < 5; i++ {
		card1Value := SignToNumber(string(hand1[i]))
		card2Value := SignToNumber(string(hand2[i]))

		if card1Value > card2Value {
			return 1
		} else if card1Value < card2Value {
			return -1
		} else {
			continue
		}
	}

	return 0
}

func CompareHands2(hand1 string, hand2 string) int {
	score1 := GetHandScore(ReplaceJokersWithBetter(hand1))
	score2 := GetHandScore(ReplaceJokersWithBetter(hand2))

	// fmt.Printf("Hand1: %s, Hand2: %s\n", hand1, hand2)

	if score1 > score2 {
		return 1
	}

	if score1 < score2 {
		return -1
	}

	for i := 0; i < 5; i++ {
		card1Value := SignToNumber(string(hand1[i]))
		card2Value := SignToNumber(string(hand2[i]))

		if card1Value > card2Value {
			return 1
		} else if card1Value < card2Value {
			return -1
		} else {
			continue
		}
	}

	return 0
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// Joker has to have value 11 istead of 1 for part 1!
func GetAnswer(lines chan string) int {
	players := []PlayerHand{}

	for line := range lines {
		hand := line[0:5]
		weight, err := strconv.Atoi(line[6:])
		if err != nil {
			fmt.Println(err)
			return 1
		}
		players = append(players, PlayerHand{hand, weight})
	}

	sort.Slice(players, func(i, j int) bool {
		return CompareHands(players[i].Hand, players[j].Hand) == -1
	})

	result := 0
	for i, hand := range players {
		result += hand.Weight * (i + 1)
	}

	return result
}

func GetAnswer2(lines chan string) int {
	players := []PlayerHand{}

	for line := range lines {
		hand := line[0:5]
		weight, err := strconv.Atoi(line[6:])
		if err != nil {
			fmt.Println(err)
			return 1
		}
		players = append(players, PlayerHand{hand, weight})
	}

	sort.Slice(players, func(i, j int) bool {
		return CompareHands2(players[i].Hand, players[j].Hand) == -1
	})

	result := 0
	for i, hand := range players {
		result += hand.Weight * (i + 1)
	}

	return result
}
