package day3

import "testing"

func TestGearsPart1(t *testing.T) {
	res1 := GetEngineSum(MockIterateLines)
	if res1 != 4361 {
		t.Errorf("Expected 4361, got %d", res1)
	}
}

func TestGearsPart2(t *testing.T) {
	res2 := GetEngineSum2(MockIterateLines)
	if res2 != 467835 {
		t.Errorf("Expected 467835, got %d", res2)
	}
}
