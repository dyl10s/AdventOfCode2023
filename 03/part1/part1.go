package main

import (
	"errors"
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

	total := 0

	for rowIndex, row := range data {
		for colIndex := 0; colIndex < len(row); colIndex++ {
			if isNumber(row[colIndex]) {
				numberString := row[colIndex]
				numberLength := 1

				for len(row) > colIndex+numberLength && isNumber(row[colIndex+numberLength]) {
					numberString += row[colIndex+numberLength]
					numberLength++
				}

				if isPartNumber(rowIndex, colIndex, numberLength, data) {
					fmt.Println("Found part number:", numberString, " at row:", rowIndex, "col:", colIndex, "with length of", numberLength)
					val, err := strconv.Atoi(numberString)
					if err != nil {
						panic("Number was not a number" + numberString)
					}

					total += val
				}

				colIndex += numberLength
			}
		}
	}

	fmt.Println("Sum of part numbers:", total)
}

func isPartNumber(startRow int, startCol int, length int, data [][]string) bool {
	for charNumber := -1; charNumber < length+1; charNumber++ {
		sameRow, sameRowErr := safeIndexData(startRow, startCol+charNumber, data)
		if sameRowErr == nil && isSymbol(sameRow) {
			return true
		}

		if startRow > 0 {
			topRow, topRowErr := safeIndexData(startRow-1, startCol+charNumber, data)
			if topRowErr == nil && isSymbol(topRow) {
				return true
			}
		}

		if startRow+1 < len(data) {
			bottomRow, bottomRowErr := safeIndexData(startRow+1, startCol+charNumber, data)
			if bottomRowErr == nil && isSymbol(bottomRow) {
				return true
			}
		}
	}

	return false
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

func isSymbol(data string) bool {
	if !isNumber(data) && data != "." {
		return true
	}

	return false
}
