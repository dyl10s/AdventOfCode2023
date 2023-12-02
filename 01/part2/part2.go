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

	numMap := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	for _, e := range strings.Split(stringContent, "\n") {
		// Disreguard blanks
		if e == "" {
			continue
		}

		firstNum := -1
		lastNum := -1

		for i, char := range e {
			foundNumber, err := strconv.Atoi(string(char))

			if err != nil {
				foundStringMatch := false

				// Replace any number strings with their number values
				for key, val := range numMap {
					if len(e) >= i+len(key) && key == e[i:i+len(key)] {
						foundNumber = val
						foundStringMatch = true
						break
					}
				}

				if foundStringMatch == false {
					continue
				}
			}

			fmt.Println("Found number", foundNumber)

			if firstNum == -1 {
				firstNum = foundNumber
			}

			lastNum = foundNumber
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
