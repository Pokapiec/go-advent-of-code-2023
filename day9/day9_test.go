package day9

import (
	"advent/utils"
	"testing"
)

func TestPredictNext(t *testing.T) {
	result := SumUpAllHistoryNumbers(utils.MockIterateLines([]string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}))

	if result != 114 {
		t.Errorf("PredictNextNumber expected 114, got %d", result)
	}

	//Part 2
	result = SumUpAllHistoryNumbersPart2(utils.MockIterateLines([]string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}))

	if result != 2 {
		t.Errorf("PredictNextNumber expected 2, got %d", result)
	}
}
