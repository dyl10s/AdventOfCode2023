package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	cachedCardResults := map[int]int{}

	content, err := ioutil.ReadFile("./data/part1.txt")
	if err != nil {
		panic(err)
	}

	stringContent := string(content)
	totalProcessed := 0

	winningNumbers := []int{}
	gameLines := strings.Split(stringContent, "\n")

	for x := len(gameLines) - 1; x >= 0; x-- {
		gameData := gameLines[x]

		// Disreguard blanks
		if gameData == "" {
			continue
		}

		numberData := strings.Split(strings.Split(gameData, ":")[1], "|")
		cardNumberSplit := strings.Split(strings.Split(gameData, ":")[0], " ")
		cardNumber, err := strconv.Atoi(cardNumberSplit[len(cardNumberSplit)-1])

		if err != nil {
			panic("Error parsing card number")
		}

		if len(numberData) != 2 {
			panic("Issue parsing the numbers for " + gameData)
		}

		winningNumbers = mapNumbersToArray(numberData[0])

		totalCardValue := 1
		winners := 0

		for _, myNumber := range mapNumbersToArray(numberData[1]) {
			if contains(winningNumbers, myNumber) {
				winners += 1
			}
		}

		for curCard := 1; curCard <= winners; curCard++ {
			totalCardValue += cachedCardResults[cardNumber+curCard]
		}

		cachedCardResults[cardNumber] = totalCardValue
		fmt.Println(cardNumber, "is value", totalCardValue, "and won", winners)
		totalProcessed += totalCardValue
	}

	fmt.Println(totalProcessed)
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
