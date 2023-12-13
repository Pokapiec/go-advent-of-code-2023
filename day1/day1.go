package day1

import (
	"fmt"
	"os"
	"strconv"
)

func CalibrationValue1(line string) int {
	var first, second int
	firstSet := false

	for _, c := range line {
		letter := string(c)
		if value, err := strconv.Atoi(letter); err == nil {
			if !firstSet {
				first = value
				firstSet = true
			}
			second = value
		}
	}

	return first*10 + second
}

func ExtractDigitFromText(line string, index int) int {
	digitPerLetter := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	res := ""

	for i := index; i < len(line); i++ {
		res += string(line[i])
		if value, ok := digitPerLetter[res]; ok {
			return value
		}

	}
	return -1
}

func CalibrationValue2(line string) int {
	var first, second int
	firstSet := false

	for i, c := range line {
		letter := string(c)
		if value, err := strconv.Atoi(letter); err == nil {
			if !firstSet {
				first = value
				firstSet = true
			}
			second = value
		} else if value := ExtractDigitFromText(line, i); value != -1 {
			if !firstSet {
				first = value
				firstSet = true
			}
			second = value
		}
	}

	return first*10 + second
}

func GetAnswer() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var line string
	sum := 0
	for {
		_, err := fmt.Fscanf(file, "%s", &line)
		if err != nil {
			break
		}
		sum += CalibrationValue2(line)
	}

	fmt.Println(sum)
}
