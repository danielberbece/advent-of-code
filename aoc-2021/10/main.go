package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	input, _ := os.Open("./input")
	scanner := bufio.NewScanner(input)

	syntaxScore, autocompleteScores := 0, make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		corrupted, points, remainingStack := checkBrackets(line)
		if corrupted {
			syntaxScore += points
		} else {
			autocompleteScores = append(autocompleteScores, completeLine(remainingStack))
		}
	}
	fmt.Println("Part 1:", syntaxScore)

	sort.Ints(autocompleteScores)
	fmt.Println("Part 2:", autocompleteScores[len(autocompleteScores)/2])
}

func completeLine(stack []rune) int {
	points := 0
	for i := len(stack) - 1; i >= 0; i-- {
		points *= 5
		if stack[i] == '(' {
			points += 1
		} else if stack[i] == '[' {
			points += 2
		} else if stack[i] == '{' {
			points += 3
		} else if stack[i] == '<' {
			points += 4
		}
	}

	return points
}

func checkBrackets(line string) (bool, int, []rune) {
	stack := make([]rune, 0)
	for _, r := range line {
		top := ' '
		if len(stack) > 0 {
			top = stack[len(stack)-1]
		}
		if r == '}' {
			if top == '{' {
				n := len(stack) - 1
				stack = stack[:n]
			} else {
				return true, 1197, stack
			}
		} else if r == ']' {
			if top == '[' {
				n := len(stack) - 1
				stack = stack[:n]
			} else {
				return true, 57, stack
			}
		} else if r == '>' {
			if top == '<' {
				n := len(stack) - 1
				stack = stack[:n]
			} else {
				return true, 25137, stack
			}
		} else if r == ')' {
			if top == '(' {
				n := len(stack) - 1
				stack = stack[:n]
			} else {
				return true, 3, stack
			}
		} else {
			stack = append(stack, r)
		}
	}

	return false, 0, stack
}
