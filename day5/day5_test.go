package day5

import "testing"

func TestWinningPointsPart2(t *testing.T) {
	result := GetAnswer(MockIterateLines)

	if result != 35 {
		t.Errorf("GetAnswer(MockIterateLines) = %v", result)
	}
}
