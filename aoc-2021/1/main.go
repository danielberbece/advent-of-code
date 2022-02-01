package main

import (
    "bufio"
    "fmt"
    //"io"
    "strconv"
    "os"
)

func main() {
	input, _ := os.Open("./input")
	scanner := bufio.NewScanner(input)
	first := true
	prev := 0
	current := 0
	cnt := 0
	depths := make([]int, 0)
	for scanner.Scan() {
		prev = current
		current, _ = strconv.Atoi(scanner.Text())
		if (!first && prev < current) {
			cnt += 1
		}
		first = false
		depths = append(depths, current)
	}
	fmt.Println("Part 1:", cnt)

	// Part 22
	cnt = 0
	for i:= 0; i < len(depths) - 3; i++ {
		sumA := depths[i] + depths[i + 1] + depths[i + 2]
		sumB := depths[i + 1] + depths[i + 2] + depths[i + 3]
		if (sumA < sumB) {
			cnt ++
		}
	}

	fmt.Println("Part 2:", cnt)
}

