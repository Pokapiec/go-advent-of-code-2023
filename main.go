package main

import (
	"advent/day9"
	"advent/utils"
	"fmt"
)

func main() {
	fmt.Println(day9.SumUpAllHistoryNumbersPart2(utils.MockIterateLines([]string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	})))

	fmt.Println(day9.SumUpAllHistoryNumbersPart2(utils.IterateLines("day9/day9_data.txt")))
}
