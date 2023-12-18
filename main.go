package main

import (
	"advent/day8"
	"advent/utils"
	"fmt"
)

func main() {
	day8Struct := day8.LRTraversal{
		FileChan: utils.IterateLines("day8/day8_data.txt"),
		// FileChan: utils.MockIterateLines([]string{
		// 	"LR",
		// 	"",
		// 	"11A = (11B, XXX)",
		// 	"11B = (XXX, 11Z)",
		// 	"11Z = (11B, XXX)",
		// 	"22A = (22B, XXX)",
		// 	"22B = (22C, 22C)",
		// 	"22C = (22Z, 22Z)",
		// 	"22Z = (22B, 22B)",
		// 	"XXX = (XXX, XXX)",
		// }),
	}

	fmt.Println(day8Struct.GetJumpsPart2())
}
