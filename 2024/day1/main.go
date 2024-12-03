package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	c, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(c), "\n")
	var col1 []int
	var col2 []int
	for i, row := range rows {
		cols := strings.Split(string(row), "   ")
		if cols[0] != "" {
			n1, err := strconv.Atoi(cols[0])
			if err != nil {
				panic(err)
			}
			col1 = append(col1, n1)
			n2, err := strconv.Atoi(cols[1])
			if err != nil {
				panic(err)
			}
			col2 = append(col2, n2)
			fmt.Println(i, n1, n2)
		}
	}

	sort.Slice(col1, func(i, j int) bool {
		return col1[i] < col1[j]
	})

	sort.Slice(col2, func(i, j int) bool {
		return col2[i] < col2[j]
	})

	var total float64
	for i := range col1 {
		total = total + math.Abs(float64(col1[i]-col2[i]))
	}
	fmt.Println("Part 1 Total: ", int(total))

	var total2 int
	for _, c1 := range col1 {
		count := 0
		for _, c2 := range col2 {
			if c1 == c2 {
				count++
			}
		}
		total2 = total2 + (c1 * count)
	}
	fmt.Println("Part 2 Total: ", int(total2))
}
