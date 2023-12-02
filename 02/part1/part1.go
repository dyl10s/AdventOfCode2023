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

	var total int

	for _, game := range strings.Split(string(content), "\n") {
		if game == "" {
			continue
		}

		idString := strings.Split(strings.Split(game, ":")[0], " ")[1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			panic("Did not find a valid game id " + idString)
		}

		shownHands := strings.Split(strings.Split(game, ":")[1], ";")

		if isGameValid(shownHands) == true {
			total += id
			fmt.Println("Game", id, "is not valid")
		} else {
			fmt.Println("Game", id, "is valid")
		}
	}

	fmt.Println("Answer is:", total)
}

func isGameValid(shownHands []string) bool {
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

	if maxBlue > 14 || maxGreen > 13 || maxRed > 12 {
		return false
	}

	return true
}
