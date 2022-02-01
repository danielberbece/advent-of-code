package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	m, folds := readInput("./input")

	for k, fold := range folds {
		m = doFold(m, fold)
		if k == 0 {
			fmt.Println("Part 1:", countPoints(m))
		}
	}

	fmt.Println("Part 2:\n")
	printMap(m)
}

func printMap(m [][]bool) {
	for _, line := range m {
		for _, point := range line {
			if point {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}

func countPoints(harta [][]bool) int {
	numPoints := 0
	for _, line := range harta {
		for _, point := range line {
			if point {
				numPoints++
			}
		}
	}

	return numPoints
}

func doFold(harta [][]bool, fold fold) [][]bool {
	newY := len(harta)
	if fold.axis == "y" {
		newY = newY / 2
	}

	newX := len(harta[0])
	if fold.axis == "x" {
		newX = newX / 2
	}
	newHarta := make([][]bool, newY)
	for i := 0; i < newY; i++ {
		newHarta[i] = make([]bool, newX)

		for j := 0; j < newX; j++ {
			if fold.axis == "y" {
				newHarta[i][j] = harta[i][j] || harta[len(harta)-i-1][j]
			} else {
				newHarta[i][j] = harta[i][j] || harta[i][len(harta[0])-j-1]
			}
		}
	}

	return newHarta
}

func readInput(path string) ([][]bool, []fold) {
	input, _ := os.Open(path)
	scanner := bufio.NewScanner(input)

	harta := make(map[point]bool)
	folds := make([]fold, 0)
	maxX, maxY := 0, 0
	inputPhase := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			inputPhase = 1
			continue
		}

		if inputPhase == 0 {
			points := strings.Split(line, ",")

			x, _ := strconv.Atoi(points[0])
			y, _ := strconv.Atoi(points[1])
			harta[point{
				x,
				y,
			}] = true

			if maxX < x {
				maxX = x
			}

			if maxY < y {
				maxY = y
			}
		} else {
			tokens := strings.Split(line, " ")
			tokens = strings.Split(tokens[2], "=")
			line, _ := strconv.Atoi(tokens[1])

			folds = append(folds, fold{axis: tokens[0], line: line})
		}
	}

	h := make([][]bool, maxY+1)
	for i := 0; i < maxY+1; i++ {
		h[i] = make([]bool, maxX+1)

		for j := 0; j < maxX+1; j++ {
			if _, present := harta[point{x: j, y: i}]; present {
				h[i][j] = true
			} else {
				h[i][j] = false
			}
		}
	}

	return h, folds
}

type point struct {
	x int
	y int
}

type fold struct {
	axis string
	line int
}
