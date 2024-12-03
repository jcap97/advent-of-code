package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.ReadFile("./example.txt")
	// file, err := os.ReadFile("./example2.txt")
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	doNots := strings.Split(string(file), "don't()")

	total := 0
	initial := true
	for _, doNot := range doNots {
		if initial {
			total = total + performValidMultiplications(doNot)
			initial = false
		} else {
			dos := strings.Split(doNot, "do()")
			if len(dos) == 1 {
				continue
			}

			total = total + performValidMultiplications(strings.Join(dos[1:], ""))
		}
	}

	fmt.Println("Total: ", total)
}

func performValidMultiplications(file string) int {
	input := strings.Split(file, "mul(")
	var multiplications [][]int
	for _, mul := range input {
		index := strings.IndexRune(mul, ')')
		if index == -1 {
			continue
		}

		numbers := strings.Split(mul[0:index], ",")
		if len(numbers) != 2 {
			continue
		}

		l, lErr := strconv.Atoi(numbers[0])
		r, rErr := strconv.Atoi(numbers[1])
		if lErr != nil || rErr != nil {
			continue
		}

		multiplications = append(multiplications, []int{l, r})
	}

	total := 0
	for _, mul := range multiplications {
		total = total + (mul[0] * mul[1])
	}

	return total
}
