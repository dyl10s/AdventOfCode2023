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

	winningNumbers := []int{}

	for _, gameData := range strings.Split(stringContent, "\n") {
		cardTotal := 0

		// Disreguard blanks
		if gameData == "" {
			continue
		}

		numberData := strings.Split(strings.Split(gameData, ":")[1], "|")

		if len(numberData) != 2 {
			panic("Issue parsing the numbers for " + gameData)
		}

		winningNumbers = mapNumbersToArray(numberData[0])

		for _, myNumber := range mapNumbersToArray(numberData[1]) {
			if contains(winningNumbers, myNumber) {
				if cardTotal == 0 {
					cardTotal = 1
				} else {
					cardTotal *= 2
				}
			}
		}

		total += cardTotal
	}

	fmt.Println(total)
}

func contains(a []int, match int) bool {
	for _, val := range a {
		if val == match {
			return true
		}
	}

	return false
}

func mapNumbersToArray(data string) []int {
	data = strings.TrimSpace(data)
	results := []int{}

	for _, num := range strings.Split(data, " ") {
		// Ignore blanks
		if num == "" || num == " " {
			continue
		}

		parsedNum, err := strconv.Atoi(num)

		if err != nil {
			panic("Error parsing number " + num + " with data: " + data)
		}

		results = append(results, parsedNum)
	}

	return results
}
