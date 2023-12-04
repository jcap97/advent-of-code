package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	var sum int
	var powers []int
	for _, line := range lines {
		numbers := strings.Split(strings.Split(string(line), ": ")[1], " | ")
		winningList := strings.Split(numbers[0], " ")
		currentList := strings.Split(numbers[1], " ")

		var power int
		for _, current := range currentList {
			for _, winning := range winningList {
				if current != "" && current == winning {
					power = power + 1
				}
			}
		}
		powers = append(powers, power)
		sum = sum + int(math.Pow(2, float64(power-1)))
	}

	fmt.Println(sum)

	instances := make([]int, len(powers))
	for i := range instances {
		instances[i] = 1
	}
	for i, p := range powers {
		for j := 0; j < p; j++ {
			for k := 0; k < instances[i]; k++ {
				if j < len(powers) {
					instances[i+j+1] = instances[i+j+1] + 1
				}
			}
		}
	}

	var total int
	for _, instance := range instances {
		total = total + instance
	}
	fmt.Println(total)
}
