package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./data/part1.txt")
	if err != nil {
		panic(err)
	}

	stringContent := string(content)
	total := 0

	for _, e := range strings.Split(stringContent, "\n") {
		// Disreguard blanks
		if e == "" {
			continue
		}

		firstNum := -1
		lastNum := -1

		for _, char := range e {
			val, err := strconv.Atoi(string(char))

			if err != nil {
				continue
			}

			fmt.Println("Found number", val)

			if firstNum == -1 {
				firstNum = val
			}

			lastNum = val
		}

		numString := strconv.Itoa(firstNum) + strconv.Itoa(lastNum)
		num, err := strconv.Atoi(numString)

		if err != nil {
			panic("Invalid number " + numString)
		}

		total += num
	}

	fmt.Print(total)
}
