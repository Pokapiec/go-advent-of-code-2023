package day8

import (
	"strings"
)

type (
	LRTraversal struct {
		FileChan chan string
	}

	Directions struct {
		Left, Right string
	}
)

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func CalculateJumpsToZ(starting, directionCombinations string, data map[string]Directions) int {
	maxDirectionIndex := len(directionCombinations) - 1
	currentDirection := 0
	jumps := 0

	current := starting
	for {
		if current[2] == 'Z' {
			break
		}

		if currentDirection > maxDirectionIndex {
			currentDirection = 0
		}

		if directionCombinations[currentDirection] == 'L' {
			current = data[current].Left
		} else {
			current = data[current].Right
		}

		currentDirection++
		jumps++
	}

	return jumps
}

func (l *LRTraversal) GetJumps() int {
	data := make(map[string]Directions)
	directionCombinations := ""

	for line := range l.FileChan {
		if directionCombinations == "" {
			directionCombinations = line
			continue
		}

		if line == "" {
			continue
		}

		source := line[0:3]
		directions := strings.Split(strings.Split(line, " = ")[1][1:9], ", ")
		data[source] = Directions{Left: directions[0], Right: directions[1]}
	}

	current := "AAA"
	maxDirectionIndex := len(directionCombinations) - 1
	currentDirection := 0
	jumps := 0

	for {
		if current == "ZZZ" {
			break
		}

		if currentDirection > maxDirectionIndex {
			currentDirection = 0
		}

		if directionCombinations[currentDirection] == 'L' {
			current = data[current].Left
		} else {
			current = data[current].Right
		}

		currentDirection++
		jumps++
	}

	return jumps
}

func (l *LRTraversal) GetJumpsPart2() int {
	data := make(map[string]Directions)
	directionCombinations := ""
	currents := []string{}

	for line := range l.FileChan {
		if directionCombinations == "" {
			directionCombinations = line
			continue
		}

		if line == "" {
			continue
		}

		source := line[0:3]
		directions := strings.Split(line[7:15], ", ")

		data[source] = Directions{Left: directions[0], Right: directions[1]}

		if source[2] == 'A' {
			currents = append(currents, source)
		}
	}

	jumps := []int{}
	for _, current := range currents {
		jumps = append(jumps, CalculateJumpsToZ(current, directionCombinations, data))
	}

	return LCM(jumps[0], jumps[1], jumps[2:]...)
}
