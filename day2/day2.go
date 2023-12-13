package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Return id of game if possible, else return 0
func IsCubeGamePossible(input string) int {
	allowedColors := map[string]int{
		"green": 13,
		"red":   12,
		"blue":  14,
	}

	data := strings.Split(input, ": ")
	gameId, err := strconv.Atoi(strings.Split(data[0], " ")[1])
	if err != nil {
		fmt.Println(err)
		return 0
	}
	cubeGames := data[1]

	// fmt.Printf("Games: %s\n", data[1])

	for _, game := range strings.Split(cubeGames, "; ") {
		for _, cube := range strings.Split(game, ", ") {
			splitedCube := strings.Split(cube, " ")
			cubeCount, err := strconv.Atoi(splitedCube[0])
			if err != nil {
				fmt.Println(err)
				return 0
			}
			cubeColor := splitedCube[1]
			restriction, ok := allowedColors[cubeColor]
			// fmt.Printf("cubeCount: %d, cubeColor: %s, restriction: %d\n", cubeCount, cubeColor, restriction)

			if !ok {
				fmt.Printf("Color %s is not allowed\n", cubeColor)
				continue
			}
			if cubeCount > restriction {
				return 0
			}
		}

	}

	return gameId
}

func MinumumCubesToPlay(input string) int {
	data := strings.Split(input, ": ")
	cubeGames := data[1]

	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, game := range strings.Split(cubeGames, "; ") {
		for _, cube := range strings.Split(game, ", ") {
			splitedCube := strings.Split(cube, " ")
			cubeCount, err := strconv.Atoi(splitedCube[0])
			if err != nil {
				fmt.Println(err)
				return 0
			}
			cubeColor := splitedCube[1]

			if cubeColor == "red" {
				if cubeCount > maxRed {
					maxRed = cubeCount
				}
			} else if cubeColor == "green" {
				if cubeCount > maxGreen {
					maxGreen = cubeCount
				}
			} else if cubeColor == "blue" {
				if cubeCount > maxBlue {
					maxBlue = cubeCount
				}
			}
		}

	}

	return maxBlue * maxGreen * maxRed
}

func GetAnswer2() {
	file, err := os.Open("day2_data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var line string
	sum := 0
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line = fileScanner.Text()
		sum += MinumumCubesToPlay(line)
	}

	fmt.Println(sum)
}
