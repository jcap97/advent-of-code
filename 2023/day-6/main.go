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

	// strTimes := strings.Fields(strings.Split(lines[0], ":")[1])
	// strDistances := strings.Fields(strings.Split(lines[1], ":")[1])
	// var (
	// 	times     []int
	// 	distances []int
	// )
	// for i := 0; i < len(strTimes); i++ {
	// 	t, _ := strconv.Atoi(strTimes[i])
	// 	d, _ := strconv.Atoi(strDistances[i])

	// 	times = append(times, t)
	// 	distances = append(distances, d)
	// }

	// moe := 1
	// for i, time := range times {
	// 	var wins []int
	// 	for j := 0; j <= time; j++ {
	// 		speed := j
	// 		travel := time - speed
	// 		distance := travel * speed

	// 		if distance > distances[i] {
	// 			wins = append(wins, distance)
	// 		}
	// 	}
	// 	moe = moe * len(wins)
	// }
	// fmt.Println(moe)

	strTime := strings.Join(strings.Fields(strings.Split(lines[0], ":")[1]), "")
	strDistance := strings.Join(strings.Fields(strings.Split(lines[1], ":")[1]), "")

	time, _ := strconv.Atoi(strTime)
	distance, _ := strconv.Atoi(strDistance)

	fmt.Println(time)
	fmt.Println(distance)

	var wins []int
	for j := 0; j <= time; j++ {
		speed := j
		travel := time - speed
		d := travel * speed

		if d > distance {
			wins = append(wins, d)
		}
	}
	fmt.Println(len(wins))
}
