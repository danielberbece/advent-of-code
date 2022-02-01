package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	arr := readInput("./input")

	// Part 1
	linearMoveCost := func(steps int) int { return steps }
	fmt.Println("Part 1:", getMinimumCost(arr, linearMoveCost))

	// Part 2
	seriesMoveCost := func(steps int) int { return steps * (steps + 1) / 2 }
	fmt.Println("Part 2:", getMinimumCost(arr, seriesMoveCost))
}

func getMinimumCost(crabPositions []int, moveCost func(int) int) int {
	maxPos := getMaxValue(crabPositions)
	d := make([][]int, len(crabPositions)+1)

	for i := 0; i < len(d); i++ {
		d[i] = make([]int, maxPos+1)
	}

	for i, pos := range crabPositions {
		for j := 0; j < len(d[i+1]); j++ {
			d[i+1][j] = d[i][j] + moveCost(abs(j-pos))
		}
	}

	minimumCost := math.MaxInt32
	for _, v := range d[len(d)-1] {
		if minimumCost > v {
			minimumCost = v
		}
	}

	return minimumCost
}

func readInput(path string) []int {
	input, _ := ioutil.ReadFile(path)
	timers := strings.Split(string(input), ",")

	arr := make([]int, 0)

	for _, timer := range timers {
		number, _ := strconv.Atoi(timer)
		arr = append(arr, number)
	}

	return arr
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func getMaxValue(arr []int) int {
	maxPos := 0
	for _, pos := range arr {
		if maxPos < pos {
			maxPos = pos
		}
	}

	return maxPos
}
