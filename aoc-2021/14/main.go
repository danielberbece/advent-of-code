package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var lastTemplateLetter int

func main() {
	polymer, rules := readInput("./input")
	part := 1
	for i := 0; i < 40; i++ {
		polymer = polymer_step(polymer, rules)

		if i == 9 || i == 39 {
			cntMostCommon := countMostCommon(polymer)
			cntLeastCommon := countLeastCommon(polymer)
			fmt.Printf("Part %d: %d\n", part, cntMostCommon-cntLeastCommon)
			part++
		}
	}
}

func countMostCommon(polymer map[[2]int]int) int {
	letter := findMostCommon(polymer)
	return count(polymer, letter)
}

func countLeastCommon(polymer map[[2]int]int) int {
	letter := findLeastCommon(polymer)
	return count(polymer, letter)
}

func findLeastCommon(polymer map[[2]int]int) int {
	m := make(map[int]int)
	minLetterCount := math.MaxInt64
	minLetter := 0
	m[lastTemplateLetter] = 1
	for pair, cnt := range polymer {
		m[pair[0]] += cnt
		m[pair[1]] += cnt
	}

	for k := range m {
		if m[k] < minLetterCount {
			minLetter = k
			minLetterCount = m[k]
		}
	}

	return minLetter
}

func findMostCommon(polymer map[[2]int]int) int {
	m := make(map[int]int)
	maxLetterCount := 0
	maxLetter := 0
	m[lastTemplateLetter] = 1
	for pair, cnt := range polymer {
		m[pair[0]] += cnt
		m[pair[1]] += cnt
		if m[pair[0]] > maxLetterCount {
			maxLetter = pair[0]
			maxLetterCount = m[pair[0]]
		}
		if m[pair[1]] > maxLetterCount {
			maxLetter = pair[1]
			maxLetterCount = m[pair[1]]
		}
	}

	return maxLetter
}

func count(polymer map[[2]int]int, letter int) int {
	cnt := 0
	for pair, cntPair := range polymer {
		if pair[0] == letter {
			cnt += cntPair
		}
	}
	if letter == lastTemplateLetter {
		cnt++
	}

	return cnt
}

func polymer_step(polymer map[[2]int]int, rules map[[2]int]int) map[[2]int]int {

	new_polymer := make(map[[2]int]int, 0)
	for pair, cnt := range polymer {
		new_letter := rules[pair]
		new_polymer[[2]int{pair[0], new_letter}] += cnt
		new_polymer[[2]int{new_letter, pair[1]}] += cnt
	}

	return new_polymer
}

func readInput(path string) (map[[2]int]int, map[[2]int]int) {
	input, _ := os.Open(path)
	scanner := bufio.NewScanner(input)

	rules := make(map[[2]int]int)
	inputPhase := 0
	polymer := make(map[[2]int]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			inputPhase = 1
			continue
		}

		if inputPhase == 0 {
			for i := 0; i < len(line)-1; i++ {
				polymer[[2]int{int(line[i]), int(line[i+1])}]++
			}
			lastTemplateLetter = int(line[len(line)-1])
		} else {
			tokens := strings.Split(line, " ")
			r := [2]int{int(tokens[0][0]), int(tokens[0][1])}
			rules[r] = int(tokens[2][0])
		}
	}

	return polymer, rules
}
