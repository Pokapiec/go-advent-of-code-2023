package utils

import (
	"bufio"
	"fmt"
	"os"
)

func IterateLines(filePath string) chan string {
	ch := make(chan string)
	go func() {
		file, err := os.Open(filePath)
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

func MockIterateLines(lines []string) chan string {
	returnChanel := make(chan string)
	go func() {
		for _, line := range lines {
			returnChanel <- line
		}
		close(returnChanel)
	}()

	return returnChanel

}
