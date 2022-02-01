package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	col int
	row int
}

type card struct {
	numbersPosition map[int]position
	size            int
	rowFill         []int
	columnFill      []int
}

func main() {
	draws, cards := readInput("./input")
	part1(draws, cards)
	// part2()
}

func part1(draws []int, cards []card) {
	won := make(map[int]int)

	for i, draw := range draws {
		for j, card := range cards {
			_, present := won[j]
			if present {
				continue
			}
			position, present := card.numbersPosition[draw]
			if present {
				card.columnFill[position.col]++
				card.rowFill[position.row]++
			}
			if card.checkWin() {
				score := card.computeScore(draws[:i+1], draw)
				if len(won) == 0 {
					fmt.Println("First card win score:", score)
				} else if len(won) == len(cards)-1 {
					fmt.Println("Last card win score:", score)
				}

				won[j] = 1
			}
		}
	}
}

// func part2() {
// 	fmt.Println("Part 2")
// }

func readInput(path string) ([]int, []card) {
	input, _ := os.Open(path)
	scanner := bufio.NewScanner(input)

	lines := make([]string, 0)
	cards := make([]card, 0)

	i := 0
	draws := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			tokens := strings.Split(line, ",")
			for _, token := range tokens {
				number, _ := strconv.Atoi(token)
				draws = append(draws, number)
			}
			scanner.Scan()
		} else if line == "" {
			cards = append(cards, newCard(lines))
			lines = make([]string, 0)
		} else {
			lines = append(lines, line)
		}

		i++
	}
	cards = append(cards, newCard(lines))

	return draws, cards
}

func newCard(lines []string) card {
	size := len(lines)

	matrix := make(map[int]position)
	for i := 0; i < size; i++ {
		lines[i] = lines[i] + " "
		for j := 0; j < size; j++ {
			str := strings.TrimSpace(lines[i][j*3 : (j+1)*3])
			number, _ := strconv.Atoi(str)
			matrix[number] = position{
				col: j,
				row: i,
			}
		}
	}

	return card{
		size:            size,
		columnFill:      make([]int, size),
		rowFill:         make([]int, size),
		numbersPosition: matrix,
	}
}

func (c card) checkWin() bool {
	for _, count := range c.columnFill {
		if count == c.size {
			return true
		}
	}

	for _, count := range c.rowFill {
		if count == c.size {
			return true
		}
	}
	return false
}

func (c card) computeScore(draws []int, finalDraw int) int {
	score := 0

	for num := range c.numbersPosition {
		if !contains(draws, num) {
			score += num
		}
	}

	return score * finalDraw
}

func contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
