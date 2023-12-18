package day9

import (
	"strconv"
	"strings"
)

func StringToInts(line string) []int {
	ints := []int{}
	for _, char := range strings.Split(line, " ") {
		value, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}
		ints = append(ints, value)
	}
	return ints

}

func PredictNextNumber(ints []int) int {
	allData := [][]int{}

	allData = append(allData, ints)

	currentRow := ints
	tempRow := []int{}
	for {
		prevNum := currentRow[0]

		areAllZero := true
		for i := 1; i < len(currentRow); i++ {
			value := currentRow[i] - prevNum
			if value != 0 {
				areAllZero = false
			}
			tempRow = append(tempRow, value)
			prevNum = currentRow[i]
		}

		allData = append(allData, tempRow)

		if areAllZero {
			break
		}

		currentRow = tempRow
		tempRow = []int{}
	}

	// fmt.Println(allData)

	lastPrediction := 0
	for i := len(allData) - 2; i >= 0; i-- {
		lastPrediction += allData[i][len(allData[i])-1]
	}

	return lastPrediction
}

func SumUpAllHistoryNumbers(lines chan string) int {
	// for line := range lines {
	// 	fmt.Println(StringToInts(line))
	// }

	result := 0
	for row := range lines {
		result += PredictNextNumber(StringToInts(row))
	}

	return result
}

func PredictPreviousNumber(ints []int) int {
	allData := [][]int{}

	allData = append(allData, ints)

	currentRow := ints
	tempRow := []int{}
	for {
		prevNum := currentRow[0]

		areAllZero := true
		for i := 1; i < len(currentRow); i++ {
			value := currentRow[i] - prevNum
			if value != 0 {
				areAllZero = false
			}
			tempRow = append(tempRow, value)
			prevNum = currentRow[i]
		}

		allData = append(allData, tempRow)

		if areAllZero {
			break
		}

		currentRow = tempRow
		tempRow = []int{}
	}

	// fmt.Println(allData)

	lastPrediction := 0
	for i := len(allData) - 2; i >= 0; i-- {
		lastPrediction = allData[i][0] - lastPrediction
	}

	return lastPrediction
}

func SumUpAllHistoryNumbersPart2(lines chan string) int {
	// for line := range lines {
	// 	fmt.Println(StringToInts(line))
	// }

	result := 0
	for row := range lines {
		result += PredictPreviousNumber(StringToInts(row))
	}

	return result
}
