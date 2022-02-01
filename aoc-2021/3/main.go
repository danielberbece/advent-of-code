package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("./input")
	// input, _ := os.Open("./test")
	scanner := bufio.NewScanner(input)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	ones := countOnes(lines)
	total := len(lines)

	// Find gamma & epsilon values
	gamma := 0
	for i, one := range ones {
		if one > total-one {
			gamma += (1 << (len(ones) - i - 1))
		}
	}

	epsilon := (1 << len(ones)) - 1 - gamma

	fmt.Println("Part 1:", gamma*epsilon)
}

func part2(lines []string) {
	oxygen_rating := make([]string, len(lines))
	carbon_rating := make([]string, len(lines))
	copy(oxygen_rating, lines)
	copy(carbon_rating, lines)

	for i := 0; len(oxygen_rating) > 1; i++ {
		mostCommon := getMostCommon(oxygen_rating)
		n := 0
		for j := 0; j < len(oxygen_rating); j++ {
			if (oxygen_rating[j][i] == '1' && mostCommon[i] == 1) || (oxygen_rating[j][i] == '0' && mostCommon[i] == 0) {
				oxygen_rating[n] = oxygen_rating[j]
				n++
			}
		}

		oxygen_rating = oxygen_rating[:n]
	}

	oxygen := getValue(oxygen_rating[0])

	for i := 0; len(carbon_rating) > 1; i++ {
		mostCommon := getMostCommon(carbon_rating)
		n := 0

		for j := 0; j < len(carbon_rating); j++ {
			if (carbon_rating[j][i] == '1' && mostCommon[i] == 0) || (carbon_rating[j][i] == '0' && mostCommon[i] == 1) {
				carbon_rating[n] = carbon_rating[j]
				n++
			}
		}

		carbon_rating = carbon_rating[:n]
	}
	carbon := getValue(carbon_rating[0])

	fmt.Println("Part 2: ", oxygen*carbon)
}

func getValue(bitsArray string) int {
	value := 0
	for i, one := range bitsArray {
		if one == '1' {
			value += (1 << (len(bitsArray) - i - 1))
		}
	}

	return value
}

func getMostCommon(lines []string) []int {
	ones := countOnes(lines)
	total := len(lines)

	for i := 0; i < len(ones); i++ {
		if ones[i] >= total-ones[i] {
			ones[i] = 1
		} else {
			ones[i] = 0
		}
	}

	return ones
}

func countOnes(lines []string) []int {
	ones := make([]int, len(lines[0]))

	for _, line := range lines {
		for i, bit := range line {
			if bit == '1' {
				ones[i]++
			}
		}
	}

	return ones
}
