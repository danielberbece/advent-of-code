package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type display struct {
	digits []string
	values []string
}

func main() {
	displays := readInput("./input")
	values := findValues(displays)
	sum := 0
	for _, v := range values {
		sum += v
	}
	fmt.Println("Part 2", sum)
}

func findValues(displays []display) (values []int) {
	values = make([]int, 0)
	for _, display := range displays {
		mapping := make(map[string]int)
		revMapping := make(map[int]string)

		// Find simple digits first
		for _, v := range display.digits {
			v = SortString(v)
			if len(v) == 2 {
				mapping[v] = 1
				revMapping[1] = v
			} else if len(v) == 3 {
				mapping[v] = 7
				revMapping[3] = v
			} else if len(v) == 4 {
				mapping[v] = 4
				revMapping[4] = v
			} else if len(v) == 7 {
				mapping[v] = 8
				revMapping[8] = v
			}
		}

		for _, v := range display.digits {
			v = SortString(v)
			if len(v) == 6 {
				if contains(v, revMapping[4]) {
					mapping[v] = 9
				} else if contains(v, revMapping[1]) {
					mapping[v] = 0
				} else {
					mapping[v] = 6
					revMapping[6] = v
				}
			}
		}

		for _, v := range display.digits {
			v = SortString(v)
			if len(v) == 5 {
				if contains(v, revMapping[1]) {
					mapping[v] = 3
				} else if contains(revMapping[6], v) {
					mapping[v] = 5
				} else {
					mapping[v] = 2
				}
			}
		}

		value := 0
		for i, v := range display.values {
			v = SortString(v)
			value += int(math.Pow10(3-i)) * mapping[v]
		}
		values = append(values, value)
	}

	return
}

func contains(str string, toContain string) bool {
	for _, c := range toContain {
		contains := false
		for _, sc := range str {
			if sc == c {
				contains = true
			}
		}

		if !contains {
			return false
		}
	}

	return true
}

func readInput(path string) (displays []display) {
	input, _ := os.Open(path)
	scanner := bufio.NewScanner(input)

	easyDigitsCount := 0
	displays = make([]display, 0)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " | ")
		digits := strings.Split(tokens[0], " ")
		values := strings.Split(tokens[1], " ")
		displays = append(displays, display{digits, values})

		for _, v := range values {
			if len(v) == 2 || len(v) == 3 || len(v) == 4 || len(v) == 7 {
				easyDigitsCount++
			}
		}
	}

	fmt.Println("Part 1", easyDigitsCount)

	return
}
