package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	n := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	c, err := os.ReadFile("./calibration.txt")
	if err != nil {
		log.Fatal(err)
	}

	list := strings.Split(string(c), "\n")
	var sum int
	for _, str := range list {
		f := getFirstDigit(str, n)
		l := getLastDigit(str, n)

		fmt.Println(f, l)
		sum = sum + (f * 10) + l
	}
	fmt.Println(sum)
}

func getFirstDigit(str string, n map[string]int) int {
	for i := 0; i < len(str); i++ {
		buff := ""
		for j := i; j < len(str); j++ {
			buff = buff + string(str[j])
			_, exist := n[buff]
			if exist {
				return n[buff]
			}
		}
	}

	return 0
}

func getLastDigit(str string, n map[string]int) int {
	for i := len(str) - 1; i >= 0; i-- {
		buff := ""
		for j := i; j < len(str); j++ {
			buff = buff + string(str[j])
			_, exist := n[buff]
			if exist {
				return n[buff]
			}
		}
	}

	return 0
}
