package main

import (
	"advent/day7"
	"advent/utils"
	"fmt"
)

func main() {
	// fmt.Println("Hello, world!")
	fmt.Println(day7.GetAnswer2(utils.IterateLines("day7/day7_data.txt")))
	fmt.Println(day7.GetAnswer2(utils.MockIterateLines([]string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	})))
}
