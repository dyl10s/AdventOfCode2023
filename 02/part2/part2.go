package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./data/part2.txt")
	if err != nil {
		panic(err)
	}

	var total int

	for _, game := range strings.Split(string(content), "\n") {
		if game == "" {
			continue
		}

		idString := strings.Split(strings.Split(game, ":")[0], " ")[1]

		shownHands := strings.Split(strings.Split(game, ":")[1], ";")
		gamePower := getGamePower(shownHands)

		fmt.Println(idString, ": ", gamePower)

		total += gamePower
	}

	fmt.Println("Answer is:", total)
}

func getGamePower(shownHands []string) int {
	var maxGreen, maxBlue, maxRed int

	for _, shownHand := range shownHands {
		for _, color := range strings.Split(shownHand, ",") {
			color = strings.TrimSpace(color)
			colorSplit := strings.Split(color, " ")

			colorName := colorSplit[1]
			amount, err := strconv.Atoi(colorSplit[0])

			if err != nil {
				panic("Invalid color amount " + colorSplit[0])
			}

			if colorName == "green" {
				if amount > maxGreen {
					maxGreen = amount
				}
			}

			if colorName == "blue" {
				if amount > maxBlue {
					maxBlue = amount
				}
			}

			if colorName == "red" {
				if amount > maxRed {
					maxRed = amount
				}
			}
		}
	}

	return maxBlue * maxGreen * maxRed
}
