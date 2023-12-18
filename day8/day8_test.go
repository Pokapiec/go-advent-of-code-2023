package day8

import (
	"advent/utils"
	"testing"
)

func TestCardsGame(t *testing.T) {
	day8Struct := LRTraversal{
		// FileChan: utils.IterateLines("day8/day8_data.txt"),
		FileChan: utils.MockIterateLines([]string{
			"RL",
			"",
			"AAA = (BBB, CCC)",
			"BBB = (DDD, EEE)",
			"CCC = (ZZZ, GGG)",
			"DDD = (DDD, DDD)",
			"EEE = (EEE, EEE)",
			"GGG = (GGG, GGG)",
			"ZZZ = (ZZZ, ZZZ)",
		}),
	}

	result := day8Struct.GetJumps()

	if result != 2 {
		t.Errorf("GetJumps expected 2, got %d", result)
	}

	day8Struct = LRTraversal{
		// FileChan: utils.IterateLines("day8/day8_data.txt"),
		FileChan: utils.MockIterateLines([]string{
			"LLR",
			"",
			"AAA = (BBB, BBB)",
			"BBB = (AAA, ZZZ)",
			"ZZZ = (ZZZ, ZZZ)",
		}),
	}

	result = day8Struct.GetJumps()

	if result != 6 {
		t.Errorf("Othe GetJumps expected 6, got %d", result)
	}

	// 	data:
	// 	LR

	// 11A = (11B, XXX)
	// 11B = (XXX, 11Z)
	// 11Z = (11B, XXX)
	// 22A = (22B, XXX)
	// 22B = (22C, 22C)
	// 22C = (22Z, 22Z)
	// 22Z = (22B, 22B)
	// XXX = (XXX, XXX)
	day8Struct = LRTraversal{
		FileChan: utils.MockIterateLines([]string{
			"LR",
			"",
			"11A = (11B, XXX)",
			"11B = (XXX, 11Z)",
			"11Z = (11B, XXX)",
			"22A = (22B, XXX)",
			"22B = (22C, 22C)",
			"22C = (22Z, 22Z)",
			"22Z = (22B, 22B)",
			"XXX = (XXX, XXX)",
		}),
	}

	result = day8Struct.GetJumpsPart2()

	if result != 6 {
		t.Errorf("GetJumpsPart2 expected 2, got %d", result)
	}
}
