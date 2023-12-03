package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	gearMap := map[string][]int{}

	content, err := ioutil.ReadFile("./data/part2.txt")
	if err != nil {
		panic(err)
	}

	var data [][]string

	// Parse the information
	for _, row := range strings.Split(string(content), "\n") {
		if row == "" {
			continue
		}

		rowData := []string{}

		for _, character := range row {
			rowData = append(rowData, string(character))
		}

		data = append(data, rowData)
	}

	for rowIndex, row := range data {
		for colIndex := 0; colIndex < len(row); colIndex++ {
			if isNumber(row[colIndex]) {
				numberString := row[colIndex]
				numberLength := 1

				for len(row) > colIndex+numberLength && isNumber(row[colIndex+numberLength]) {
					numberString += row[colIndex+numberLength]
					numberLength++
				}

				val, err := strconv.Atoi(numberString)
				if err != nil {
					panic("Number was not a number" + numberString)
				}

				checkForGear(rowIndex, colIndex, numberLength, data, val, gearMap)
				colIndex += numberLength
			}
		}
	}

	total := 0

	// Loop over and add up the gear map values
	for _, value := range gearMap {
		if len(value) == 2 {
			total += value[0] * value[1]
		}
	}

	fmt.Println("The total gear ratio is:", total)
}

func checkForGear(startRow int, startCol int, length int, data [][]string, number int, gearMap map[string][]int) bool {
	for charNumber := -1; charNumber < length+1; charNumber++ {

		sameRow, sameRowErr := safeIndexData(startRow, startCol+charNumber, data)
		if sameRowErr == nil && isGear(sameRow) {
			addToGearMap(gearMap, startRow, startCol+charNumber, number)
		}

		if startRow > 0 {
			topRow, topRowErr := safeIndexData(startRow-1, startCol+charNumber, data)
			if topRowErr == nil && isGear(topRow) {
				addToGearMap(gearMap, startRow-1, startCol+charNumber, number)
			}
		}

		if startRow+1 < len(data) {
			bottomRow, bottomRowErr := safeIndexData(startRow+1, startCol+charNumber, data)
			if bottomRowErr == nil && isGear(bottomRow) {
				addToGearMap(gearMap, startRow+1, startCol+charNumber, number)
			}
		}
	}

	return false
}

func addToGearMap(gearMap map[string][]int, row int, col int, val int) {
	gearItem := gearMap[strconv.Itoa(row)+","+strconv.Itoa(col)]
	gearItem = append(gearItem, val)
	gearMap[strconv.Itoa(row)+","+strconv.Itoa(col)] = gearItem
}

func safeIndexData(row int, col int, data [][]string) (string, error) {
	if row > 0 && row < len(data) {
		if col > 0 && col < len(data[row]) {
			return data[row][col], nil
		}
	}

	return "", errors.New("Out of bounds index")
}

func isNumber(data string) bool {
	_, err := strconv.Atoi(data)

	if err != nil {
		return false
	}

	return true
}

func isGear(data string) bool {
	if data == "*" {
		return true
	}

	return false
}
