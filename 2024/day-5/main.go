package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("./example.txt")
	// file, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(file), "\n")
	rows = rows[0 : len(rows)-1]

	index := slices.Index(rows, "")

	total := 0
	for _, updates := range rows[index+1:] {
		valid := true
		pages := strings.Split(updates, ",")
		for _, pageOrdering := range rows[:index] {
			order := strings.Split(pageOrdering, "|")
			for i, page := range pages {
				if page == order[0] {
					j := slices.Index(pages, order[1])
					if j != -1 && i > j {
						valid = false
					}
				}
			}
		}

		if valid {
			val, _ := strconv.Atoi(pages[(len(pages)-1)/2])
			fmt.Println(val)
			total += val
		}
	}

	fmt.Println("Part 1 Total: ", total)
}
