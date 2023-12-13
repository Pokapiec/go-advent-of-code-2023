package day1

import (
	"testing"
)

func TestCalibrationValue1(t *testing.T) {

	var tests = []struct {
		input string
		want  int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, test := range tests {
		if got := CalibrationValue1(test.input); got != test.want {
			t.Errorf("CalibrationValue(%q) = %v", test.input, got)
		}
	}

}

func TestCalibrationValue2(t *testing.T) {

	var tests = []struct {
		input string
		want  int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}

	for _, test := range tests {
		if got := CalibrationValue2(test.input); got != test.want {
			t.Errorf("CalibrationValue(%q) = %v", test.input, got)
		}
	}

}
