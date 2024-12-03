package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// config := map[string]int{
	// 	"red":   12,
	// 	"green": 13,
	// 	"blue":  14,
	// }

	games := strings.Split(string(input), "\n")
	var sum int
	// for _, game := range games {
	// 	str := strings.Split(game, ":")
	// 	n := strings.Split(str[0], " ")[1]
	// 	val, err := strconv.Atoi(n)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if isGameValid(str[1], config) {
	// 		sum = sum + val
	// 	}
	// }

	for _, game := range games {
		str := strings.Split(game, ":")
		if err != nil {
			log.Fatal(err)
		}
		sum = sum + gamePowerSet(str[1])
	}

	fmt.Println(sum)
}

func isGameValid(str string, config map[string]int) bool {
	rounds := strings.Split(str, ";")
	for _, round := range rounds {
		colors := strings.Split(round, ",")
		for _, color := range colors {
			val := strings.Split(color, " ")
			n, err := strconv.Atoi(val[1])
			if err != nil {
				log.Fatal(err)
			}

			if n > config[val[2]] {
				return false
			}
		}
	}
	return true
}

func gamePowerSet(str string) int {
	rounds := strings.Split(str, ";")
	configColors := map[string]int{
		"red":   1,
		"green": 1,
		"blue":  1,
	}

	for _, round := range rounds {
		colors := strings.Split(round, ",")
		for _, color := range colors {
			val := strings.Split(color, " ")
			n, err := strconv.Atoi(val[1])
			if err != nil {
				log.Fatal(err)
			}
			if n > configColors[val[2]] {
				configColors[val[2]] = n
			}
		}
	}

	power := 1
	for _, val := range configColors {
		power = power * val
	}

	return power
}
