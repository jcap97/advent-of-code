package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// file, err := os.ReadFile("./example.txt")
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(file), "\n")
	rows = rows[0 : len(rows)-1]
	for i := 0; i < 4; i++ {
		padding := ""
		for j := 0; j < len(rows[0]); j++ {
			padding += "#"
		}

		rows = append([]string{padding}, rows...)
		rows = append(rows, padding)
	}

	totalPart1 := 0
	totalPart2 := 0
	var puzzle []string
	for _, row := range rows {
		puzzle = append(puzzle, "######"+row+"#####")
	}

	for i, row := range puzzle {
		for j, col := range row {
			if col == 'X' {
				// Up
				if puzzle[i-1][j] == 'M' && puzzle[i-2][j] == 'A' && puzzle[i-3][j] == 'S' {
					totalPart1 += 1
				}

				// Down
				if puzzle[i+1][j] == 'M' && puzzle[i+2][j] == 'A' && puzzle[i+3][j] == 'S' {
					totalPart1 += 1
				}

				// Left
				if puzzle[i][j-1] == 'M' && puzzle[i][j-2] == 'A' && puzzle[i][j-3] == 'S' {
					totalPart1 += 1
				}

				// Right
				if puzzle[i][j+1] == 'M' && puzzle[i][j+2] == 'A' && puzzle[i][j+3] == 'S' {
					totalPart1 += 1
				}

				// Left+Up
				if puzzle[i-1][j-1] == 'M' && puzzle[i-2][j-2] == 'A' && puzzle[i-3][j-3] == 'S' {
					totalPart1 += 1
				}

				// Right+Up
				if puzzle[i-1][j+1] == 'M' && puzzle[i-2][j+2] == 'A' && puzzle[i-3][j+3] == 'S' {
					totalPart1 += 1
				}

				// Left+Down
				if puzzle[i+1][j-1] == 'M' && puzzle[i+2][j-2] == 'A' && puzzle[i+3][j-3] == 'S' {
					totalPart1 += 1
				}

				// Right+Down
				if puzzle[i+1][j+1] == 'M' && puzzle[i+2][j+2] == 'A' && puzzle[i+3][j+3] == 'S' {
					totalPart1 += 1
				}
			}

			if col == 'A' {
				if puzzle[i-1][j+1] == 'M' && puzzle[i+1][j+1] == 'M' && puzzle[i-1][j-1] == 'S' && puzzle[i+1][j-1] == 'S' {
					totalPart2 += 1
				}

				// if puzzle[i-1][j+1] == 'S' && puzzle[i+1][j+1] == 'M' && puzzle[i-1][j-1] == 'M' && puzzle[i+1][j-1] == 'S' {
				// 	totalPart2 += 1
				// }

				if puzzle[i-1][j+1] == 'S' && puzzle[i+1][j+1] == 'S' && puzzle[i-1][j-1] == 'M' && puzzle[i+1][j-1] == 'M' {
					totalPart2 += 1
				}

				if puzzle[i-1][j+1] == 'M' && puzzle[i+1][j+1] == 'S' && puzzle[i-1][j-1] == 'M' && puzzle[i+1][j-1] == 'S' {
					totalPart2 += 1
				}

				if puzzle[i-1][j+1] == 'S' && puzzle[i+1][j+1] == 'M' && puzzle[i-1][j-1] == 'S' && puzzle[i+1][j-1] == 'M' {
					totalPart2 += 1
				}
			}
		}
	}

	fmt.Println("Total Part 1: ", totalPart1)
	fmt.Println("Total Part 2: ", totalPart2)
}
