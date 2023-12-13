package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func IterateLines() chan string {
	ch := make(chan string)
	go func() {
		file, err := os.Open("day3/day3_data.txt")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		var line string
		fileScanner := bufio.NewScanner(file)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			line = fileScanner.Text()
			ch <- line
		}

		close(ch)
	}()
	return ch
}

func MockIterateLines() chan string {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	returnChanel := make(chan string)
	go func() {
		for _, line := range lines {
			returnChanel <- line
		}
		close(returnChanel)
	}()

	return returnChanel

}

func GetEngineSum(iterLines func() chan string) int {
	lineGrid := []string{}

	for line := range iterLines() {
		lineGrid = append(lineGrid, line)
	}

	sum := 0
	currentNum := ""
	isValidated := false
	for row, line := range lineGrid {
		for col, num := range line {
			if unicode.IsDigit(num) {
				currentNum += string(num)
				if !isValidated {
					isValidated = IsCharAdjustendToAnySymbol(row, col, &lineGrid)
				}
			} else {
				if currentNum != "" {
					value, err := strconv.Atoi(currentNum)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					if isValidated {
						fmt.Println(value)
						sum += value
					}
				}
				currentNum = ""
				isValidated = false
			}
		}
	}

	if currentNum != "" {
		value, err := strconv.Atoi(currentNum)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if isValidated {
			sum += value
		}
	}

	return sum
}

func IsCharAdjustendToAnySymbol(row, col int, grid *[]string) bool {

	pointsToCheck := []struct {
		row int
		col int
	}{
		{row - 1, col - 1},
		{row - 1, col},
		{row - 1, col + 1},
		{row, col - 1},
		{row, col + 1},
		{row + 1, col - 1},
		{row + 1, col},
		{row + 1, col + 1},
	}

	for _, point := range pointsToCheck {
		if point.row < 0 || point.col < 0 {
			continue
		}
		if point.row >= len(*grid) || point.col >= len((*grid)[point.row]) {
			continue
		}
		char := (*grid)[point.row][point.col]
		if !unicode.IsDigit(rune(char)) && char != '.' {
			return true
		}
	}
	return false
}

func GetEngineSum2(iterLines func() chan string) int {
	lineGrid := []string{}

	for line := range iterLines() {
		lineGrid = append(lineGrid, line)
	}

	sum := 0
	for row, line := range lineGrid {
		for col, char := range line {
			if char == '*' {
				sum += GetGearPair(row, col, lineGrid)
			}
		}
	}

	return sum
}

func isNumNearGear(r, c, row, col int) bool {
	allowedPoints := []struct {
		row int
		col int
	}{
		{row - 1, col - 1},
		{row - 1, col},
		{row - 1, col + 1},
		{row, col - 1},
		{row, col + 1},
		{row + 1, col - 1},
		{row + 1, col},
		{row + 1, col + 1},
	}

	for _, point := range allowedPoints {
		if point.row == r && point.col == c {
			return true
		}
	}

	return false
}

func GetGearPair(row, col int, grid []string) int {

	numsFound := []int{}
	currentNum := ""

	isNumberNearGear := false

	for r, line := range grid {
		for c, char := range line {

			if r > row+1 || r < row-1 {
				continue
			}

			if unicode.IsDigit(char) {
				currentNum += string(char)
				if !isNumberNearGear {
					isNumberNearGear = isNumNearGear(r, c, row, col)
				}
			} else {
				if currentNum != "" {
					value, err := strconv.Atoi(currentNum)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					if isNumberNearGear {
						numsFound = append(numsFound, value)
					}
				}
				isNumberNearGear = false
				currentNum = ""
			}
		}
	}

	if currentNum != "" {
		value, err := strconv.Atoi(currentNum)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if isNumberNearGear {
			numsFound = append(numsFound, value)
		}
	}

	if len(numsFound) > 0 {
		fmt.Println(numsFound)

	}

	if len(numsFound) == 2 {
		return numsFound[0] * numsFound[1]
	}

	return 0
}

func Part1Executor() {
}
