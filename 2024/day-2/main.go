package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	r, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reports := strings.Split(string(r), "\n")
	initialCount, badReports := safeCount(reports)

	toleranceCount := 0
	for _, report := range badReports {
		if retryWithTolerance(report) {
			toleranceCount++
		}
	}

	fmt.Println("Part Safe Count: ", initialCount+toleranceCount)
}

func retryWithTolerance(report string) bool {
	cols := strings.Split(report, " ")
	var permutations []string
	for i := range cols {
		var permute []string
		for j, col := range cols {
			if i == j {
				continue
			}
			permute = append(permute, col)
		}
		permutations = append(permutations, strings.Join(permute, " "))
	}

	valid := false
	for _, permute := range permutations {
		if permute != "" {
			if validateReport(permute) {
				valid = true
			}
		}
	}
	return valid
}

func safeCount(reports []string) (int, []string) {
	count := 0
	var invalid []string
	for _, report := range reports {
		if report != "" {
			valid := validateReport(report)

			if valid {
				count++
			} else {
				invalid = append(invalid, report)
			}
		}
	}

	return count, invalid
}

func validateReport(report string) bool {
	valid := true
	cols := strings.Split(string(report), " ")
	var differences []int
	for i := range len(cols) - 1 {
		r, _ := strconv.Atoi(cols[i+1])
		l, _ := strconv.Atoi(cols[i])
		differences = append(differences, r-l)
	}

	for _, diff := range differences {
		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			valid = false
		}
	}

	for i := range len(differences) - 1 {
		if math.Signbit(float64(differences[i+1])) != math.Signbit(float64(differences[i])) {
			valid = false
		}
	}

	fmt.Println(report, valid)
	return valid
}
