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

	seeds := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	maps := getMaps(lines[1:])

	minLoc := 999999999999999999
	// for _, seed := range seeds {
	// 	s, err := strconv.Atoi(seed)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if loc := getLocation(s, maps); loc < minLoc {
	// 		minLoc = loc
	// 	}
	// }

	for i := 0; i < len(seeds); i = i + 2 {
		seed, _ := strconv.Atoi(seeds[i])
		rng, _ := strconv.Atoi(seeds[i+1])

		for i := seed; i < seed+rng; i++ {
			if loc := getLocation(i, maps); loc < minLoc {
				minLoc = loc
				fmt.Println("Min. Location: ", minLoc)
			}
		}
	}

	// for i := 0; i < len(seeds); i = i + 2 {
	// 	seed, _ := strconv.Atoi(seeds[i])
	// 	rng, _ := strconv.Atoi(seeds[i+1])
	// 	if loc := getLocationWithRange(seed, rng, maps); loc < minLoc {
	// 		minLoc = loc
	// 	}
	// }

	fmt.Println("Min. Location: ", minLoc)
}

func getLocationWithRange(seed int, seedRng int, maps [][][]int) int {
	loc := seed
	locRng := seedRng
	fmt.Println("Seed: ", seed)
	fmt.Print(loc, locRng)
	for _, m := range maps {
		cur := loc
		curRng := locRng

		min := cur
		max := cur + curRng - 1

		for _, path := range m {
			src := path[1]
			dst := path[0]
			rng := path[2]

			if min <= src && max >= src && max < src+rng {
				cur = dst + src - min
				curRng = max - src + 1
				fmt.Print("o-i")
			} else if min <= src && max > src+rng {
				cur = dst + src - min
				curRng = rng
				fmt.Print("o-o")
			} else if src <= min && src+rng > max {
				cur = dst + min - src
				curRng = max - min + 1
				fmt.Print("i-i")
			}

			// if loc+seedRng >= src && loc+seedRng < rng+src {
			// 	cur = dst
			// }
		}
		loc = cur
		locRng = curRng
		fmt.Print(" -> ", loc, locRng)
	}
	fmt.Println("")
	return loc
}

func getLocation(seed int, maps [][][]int) int {
	loc := seed
	// fmt.Println("Seed: ", seed)
	// fmt.Print(loc)
	for _, m := range maps {
		cur := loc
		for _, path := range m {
			src := path[1]
			dst := path[0]
			rng := path[2]

			if loc >= src && loc < rng+src {
				cur = dst + loc - src
			}
		}
		loc = cur
		// fmt.Print(" -> ", loc)
	}
	// fmt.Println("")
	return loc
}

func getMaps(lines []string) [][][]int {
	var divisions []int
	for i, line := range lines {
		if line == "" {
			divisions = append(divisions, i)
		}
	}

	var buffMaps [][][]int
	for i, division := range divisions {
		var buffLines []string
		if i == len(divisions)-1 {
			buffLines = lines[division+2:]
		} else {
			buffLines = lines[division+2 : divisions[i+1]]
		}

		var buffMap [][]int
		for _, line := range buffLines {
			numbers := strings.Split(line, " ")
			var buff []int
			for _, n := range numbers {
				val, err := strconv.Atoi(n)
				if err != nil {
					panic(err)
				}
				buff = append(buff, val)
			}
			buffMap = append(buffMap, buff)
		}

		buffMaps = append(buffMaps, buffMap)
	}

	return buffMaps
}
