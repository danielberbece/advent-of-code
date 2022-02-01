package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	polymer, rules := readInput("./test")
	for i := 0; i < 40; i++ {
		polymer = polymer_step(polymer, rules)
		fmt.Println("Step: ", i)
	}

	cntMostCommon := countMostCommon(polymer)
	cntLeastCommon := countLeastCommon(polymer)
	fmt.Println(cntMostCommon, cntLeastCommon, cntMostCommon-cntLeastCommon)
}

func countMostCommon(str string) int {
	letter := findMostCommon(str)
	return count(str, letter)
}

func countLeastCommon(str string) int {
	letter := findLeastCommon(str)
	return count(str, letter)
}

func findLeastCommon(str string) rune {
	m := make(map[rune]int)
	minLetterCount := math.MaxInt32
	minLetter := 'a'
	for _, letter := range str {
		m[letter]++
	}

	for k, v := range m {
		if v < minLetterCount {
			minLetter = k
			minLetterCount = v
		}
	}

	return minLetter
}

func findMostCommon(str string) rune {
	m := make(map[rune]int)
	maxLetterCount := 0
	maxLetter := 'a'
	for _, letter := range str {
		m[letter]++
		if m[letter] > maxLetterCount {
			maxLetter = letter
			maxLetterCount = m[letter]
		}
	}

	return maxLetter
}

func count(str string, letter rune) int {
	cnt := 0
	for _, l := range str {
		if l == letter {
			cnt++
		}
	}
	return cnt
}

func polymer_step(polymer string, rules map[string]byte) string {
	new_polymer := make([]string, 0)

	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		new_letter := rules[pair]
		new_polymer = append(new_polymer, fmt.Sprintf("%c%c", pair[0], new_letter))
	}

	new_polymer = append(new_polymer, fmt.Sprintf("%c", polymer[len(polymer)-1]))

	return strings.Join(new_polymer, "")
}

func readInput(path string) (string, map[string]byte) {
	input, _ := os.Open(path)
	scanner := bufio.NewScanner(input)

	rules := make(map[string]byte)
	// folds := make([]fold, 0)
	// maxX, maxY := 0, 0
	inputPhase := 0
	var polymer string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			inputPhase = 1
			continue
		}

		if inputPhase == 0 {
			polymer = line
		} else {
			tokens := strings.Split(line, " ")
			rules[tokens[0]] = tokens[2][0]
		}
	}

	return polymer, rules
}
