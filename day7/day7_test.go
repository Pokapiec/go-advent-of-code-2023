package day7

import (
	"advent/utils"
	"testing"
)

func TestCardsGame(t *testing.T) {
	result := GetAnswer(utils.MockIterateLines([]string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}))

	if result != 6440 {
		t.Errorf("TestCardsGame expected 6440, got %d", result)
	}

	result = GetAnswer2(utils.MockIterateLines([]string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}))

	if result != 5905 {
		t.Errorf("TestCardsGame Part2 expected 5905, got %d", result)
	}
}
