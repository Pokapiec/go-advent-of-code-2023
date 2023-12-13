package day2

import (
	"testing"
)

func TestCubeGame(t *testing.T) {

	var tests = []struct {
		input string
		want  int
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 1},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 2},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 0},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 0},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 5},
	}

	for _, test := range tests {
		if got := IsCubeGamePossible(test.input); got != test.want {
			t.Errorf("IsCubeGamePossible(%q) = %v", test.input, got)
		}
	}

}

func TestMinCubesPerGame(t *testing.T) {

	var tests = []struct {
		input string
		want  int
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 48},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 1560},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 630},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 36},
	}

	for _, test := range tests {
		if got := MinumumCubesToPlay(test.input); got != test.want {
			t.Errorf("MinumumCubesToPlay(%q) = %v", test.input, got)
		}
	}

}
