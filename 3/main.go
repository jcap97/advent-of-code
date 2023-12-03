package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	symbols := getSymbolCoordinates(lines)

	var sum int
	// for _, line := range lines {
	// groups := getLineNumbersWithCoordinates(lines)[i]
	// for _, group := range groups {
	// 	var buff string
	// 	if validNumber(line, i, group, symbols) {
	// 		for _, c := range group {
	// 			buff = buff + string(line[c])
	// 		}
	// 		n, err := strconv.Atoi(buff)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		sum = sum + n
	// 	}
	// }
	// }

	for i, symbol := range symbols {
		for _, s := range symbol {
			sum = sum + getGearRatio(lines, i, getLineNumbersWithCoordinates(lines), s)
		}
	}

	fmt.Println(sum)
}

func validNumber(line string, i int, group []int, symbols [][]int) bool {
	var config []int
	if i == 0 {
		config = []int{0, 1}
	} else if i == len(symbols)-1 {
		config = []int{-1, 0}
	} else {
		config = []int{-1, 0, 1}
	}

	searchGroup := group
	if len(line) != group[len(group)-1] {
		searchGroup = append(searchGroup, group[len(group)-1]+1)
	}
	if group[0] != 0 {
		searchGroup = append(searchGroup, group[0]-1)
	}

	for _, c := range config {
		for _, n := range searchGroup {
			for _, s := range symbols[i+c] {
				if s == n {
					return true
				}
			}
		}
	}
	return false
}

func getGearRatio(lines []string, i int, lineGroups [][][]int, symbolIndex int) int {
	var config []int
	if i == 0 {
		config = []int{0, 1}
	} else if i == len(lines)-1 {
		config = []int{-1, 0}
	} else {
		config = []int{-1, 0, 1}
	}

	var gears []int
	for _, c := range config {
		searchSymbol := []int{symbolIndex}
		if symbolIndex != 0 {
			searchSymbol = append(searchSymbol, symbolIndex-1)
		}
		if symbolIndex != len(lines[i])-1 {
			searchSymbol = append(searchSymbol, symbolIndex+1)
		}

		groups := lineGroups[i+c]

		for _, s := range searchSymbol {
			for _, group := range groups {
				for _, n := range group {
					if s == n {
						var buff string
						for _, char := range group {
							buff = buff + string(lines[i+c][char])
						}

						n, err := strconv.Atoi(buff)
						if err != nil {
							panic(err)
						}

						skip := false
						for _, gear := range gears {
							if gear == n {
								skip = true
							}
						}
						if !skip {
							gears = append(gears, n)
						}
					}
				}
			}
		}
	}

	if len(gears) == 2 {
		fmt.Println(gears[0], gears[1])
		return gears[0] * gears[1]
	} else {
		return 0
	}
}

func getLineNumbersWithCoordinates(lines []string) [][][]int {
	coordinates := make([][][]int, 0)

	for _, line := range lines {
		groups := make([][]int, 0)
		indexes := make([]int, 0)
		for j, char := range line {
			if _, err := strconv.Atoi(string(char)); err == nil {
				indexes = append(indexes, j)
			}
		}

		var group []int
		for i, index := range indexes {
			p := i + 1
			if p == len(indexes) {
				p = len(indexes) - 1
			}
			group = append(group, index)
			if index+1 != indexes[p] {
				groups = append(groups, group)
				group = nil
			}
		}
		coordinates = append(coordinates, groups)
	}

	return coordinates
}

func getSymbolCoordinates(lines []string) [][]int {
	symbolCoordinates := make([][]int, 0)
	for _, line := range lines {
		symbolIndexes := make([]int, 0)
		for j, char := range line {
			// if _, err := strconv.Atoi(string(char)); err != nil && string(char) != "." {
			// 	symbolIndexes = append(symbolIndexes, j)
			// }
			if string(char) == "*" {
				symbolIndexes = append(symbolIndexes, j)
			}
		}
		symbolCoordinates = append(symbolCoordinates, symbolIndexes)
	}
	return symbolCoordinates
}
