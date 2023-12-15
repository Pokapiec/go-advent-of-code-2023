package day6

import "testing"

func TestRacePart1(t *testing.T) {
	result := GetAnswer(MockIterateLines)

	if result != 288 {
		t.Errorf("GetAnswer(MockIterateLines) = %v", result)
	}
}

func TestRacePart2(t *testing.T) {
	result := GetAnswer2(MockIterateLines)

	if result != 71503 {
		t.Errorf("GetAnswer2(MockIterateLines) = %v", result)
	}
}
