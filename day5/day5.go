package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MockIterateLines() chan string {
	lines := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
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
		file, err := os.Open("day5/day5_data.txt")
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

func LineToInts(line string) []int {
	returnInts := []int{}
	for _, v := range strings.Split(line, " ") {
		value, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		returnInts = append(returnInts, value)
	}
	return returnInts
}

func GetAnswer(iterLines func() chan string) int {
	seedNums := []int{}
	tempSeeds := []int{}
	mappings := make(map[int]int)

	seenFirstMap := false
	// currentMap := ""

	for line := range iterLines() {
		if strings.Contains(line, "seeds: ") {
			seedNums = LineToInts(strings.Split(line, "seeds: ")[1])
		}

		if strings.Contains(line, "map:") {
			// currentMap = line
			seenFirstMap = true

			for _, seed := range seedNums {
				v, ok := mappings[seed]
				if !ok {
					tempSeeds = append(tempSeeds, seed)
				} else {
					tempSeeds = append(tempSeeds, v)
				}
			}
			seedNums = tempSeeds
			tempSeeds = tempSeeds[:0]
			mappings = make(map[int]int)

			// fmt.Printf("currentMap: %s, seeds: %v\n", currentMap, seedNums)
			continue
		}

		if seenFirstMap && line != "" {
			values := LineToInts(line)
			startSource := values[1]
			startDest := values[0]
			jumps := values[2]

			for _, v := range seedNums {
				if v >= startSource && v < startSource+jumps {
					mappings[v] = startDest + v - startSource
				}
			}
		}

	}

	for _, seed := range seedNums {
		v, ok := mappings[seed]
		if !ok {
			tempSeeds = append(tempSeeds, seed)
		} else {
			tempSeeds = append(tempSeeds, v)
		}
	}
	seedNums = tempSeeds

	minimal := seedNums[0]

	for _, seed := range seedNums {
		if seed < minimal {
			minimal = seed
		}
	}
	return minimal
}

type SeedRange struct {
	start, end int
}

// Not working yet
func GetAnswerPart2(iterLines func() chan string) int {
	seedNums := []int{}
	tempSeeds := []int{}
	mappings := make(map[int]int)

	seenFirstMap := false
	// currentMap := ""

	for line := range iterLines() {
		if strings.Contains(line, "seeds: ") {
			seedRanges := LineToInts(strings.Split(line, "seeds: ")[1])

			var num int
			var count int
			for i, seed := range seedRanges {
				if i%2 == 0 {
					num = seed
					seedNums = append(seedNums, num)
				} else {
					count = seed
					seedNums = append(seedNums, num+count)
				}
			}

			fmt.Println(seedNums)

		}

		if strings.Contains(line, "map:") {
			// currentMap = line
			seenFirstMap = true

			for _, seed := range seedNums {
				v, ok := mappings[seed]
				if !ok {
					tempSeeds = append(tempSeeds, seed)
				} else {
					tempSeeds = append(tempSeeds, v)
				}
			}
			seedNums = tempSeeds
			tempSeeds = tempSeeds[:0]
			mappings = make(map[int]int)

			// fmt.Printf("currentMap: %s, seeds: %v\n", currentMap, seedNums)
			continue
		}

		if seenFirstMap && line != "" {
			values := LineToInts(line)
			startSource := values[1]
			startDest := values[0]
			jumps := values[2]

			for _, v := range seedNums {
				if v >= startSource && v < startSource+jumps {
					mappings[v] = startDest + v - startSource
				}
			}
		}

	}

	for _, seed := range seedNums {
		v, ok := mappings[seed]
		if !ok {
			tempSeeds = append(tempSeeds, seed)
		} else {
			tempSeeds = append(tempSeeds, v)
		}
	}
	seedNums = tempSeeds
	minimal := seedNums[0]

	for _, seed := range seedNums {
		if seed < minimal {
			minimal = seed
		}
	}
	return minimal
}
