package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MockIterateLines() chan string {
	lines := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
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

func IterateLines() chan string {
	ch := make(chan string)
	go func() {
		file, err := os.Open("day6/day6_data.txt")
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

func StringToIntList(str string) []int {
	intList := []int{}
	for _, v := range strings.Split(str, " ") {
		num := strings.TrimSpace(v)
		value, err := strconv.Atoi(num)
		if err != nil {
			continue
		}

		intList = append(intList, value)
	}
	return intList
}

type Race struct {
	time, distance int
}

// Racing toys !
func GetAnswer(iterLines func() chan string) int {
	data := []Race{}
	listOfLists := [][]int{}

	for line := range iterLines() {
		splited := strings.Split(line, ":")
		nums := splited[1]
		listOfLists = append(listOfLists, StringToIntList(nums))
	}

	racesAmount := len(listOfLists[0])
	for i := 0; i < racesAmount; i++ {
		data = append(data, struct{ time, distance int }{listOfLists[0][i], listOfLists[1][i]})
	}

	fmt.Println(data)

	return CalculateRacesProduct(data)
}

func CalculateRacesProduct(input []Race) int {
	product := 1
	for _, race := range input {
		totalBetterTimes := 0
		for i := 0; i < +race.time; i++ {
			resultDistance := (race.time - i) * i
			if resultDistance > race.distance {
				totalBetterTimes++
			}
		}
		product *= totalBetterTimes
	}

	return product
}

func GetAnswer2(iterLines func() chan string) int {
	time := 0
	distance := 0

	nums := []int{}

	for line := range iterLines() {
		splited := strings.Split(line, ":")
		num := splited[1]
		num = strings.ReplaceAll(num, " ", "")
		num = strings.TrimSpace(num)
		value, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
		}
		nums = append(nums, value)
	}

	time = nums[0]
	distance = nums[1]

	return CalculateRacesProduct([]Race{{time, distance}})
}
