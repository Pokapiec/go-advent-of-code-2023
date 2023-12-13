package day4

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func StringNumsToInts(input string) []int {
	nums := []int{}
	for _, v := range strings.Split(input, " ") {
		if v == "" {
			continue
		}
		value, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			return []int{}
		}
		nums = append(nums, value)
	}
	return nums
}

func WinningPoints(input string) int {
	data := strings.Split(strings.Split(input, ": ")[1], " | ")
	winningNumbers := StringNumsToInts(data[0])
	haveNumbers := StringNumsToInts(data[1])

	won := 0
	for _, v := range haveNumbers {
		if slices.Contains(winningNumbers, v) {
			if won == 0 {
				won = 1
			} else {
				won *= 2
			}
		}
	}
	return won
}

func AmountOfMatchingCards(input string) int {
	data := strings.Split(strings.Split(input, ": ")[1], " | ")
	winningNumbers := StringNumsToInts(data[0])
	haveNumbers := StringNumsToInts(data[1])

	amount := 0
	for _, v := range haveNumbers {
		if slices.Contains(winningNumbers, v) {
			amount++
		}
	}
	return amount
}

func IterateLines() chan string {
	ch := make(chan string)
	go func() {
		file, err := os.Open("day4/day4_data.txt")
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
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
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

func GetAnswer() int {
	sum := 0
	for line := range IterateLines() {
		sum += WinningPoints(line)
	}
	return sum
}

func GetAnswerPart2(iterLines func() chan string) int {
	var amount int
	amountPerCard := make(map[int]int)

	cardNumber := 1
	for line := range iterLines() {
		v, ok := amountPerCard[cardNumber]
		if !ok {
			amountPerCard[cardNumber] = 1
		} else {
			amountPerCard[cardNumber] = v + 1
		}
		vCurr := amountPerCard[cardNumber]
		amount = AmountOfMatchingCards(line)

		// fmt.Printf("Card %d: %d\n", cardNumber, amount)
		for i := cardNumber + 1; i <= cardNumber+amount; i++ {
			v, ok := amountPerCard[i]
			if !ok {
				amountPerCard[i] = 1
			}
			amountPerCard[i] = v + vCurr
		}

		cardNumber++
	}

	sum := 0
	for _, amount := range amountPerCard {
		sum += amount
	}

	return sum
}
