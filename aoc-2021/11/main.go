package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	matrix := readInput("./input")

	flashes := 0
	for i := 0; true; i++ {
		flashes += step(matrix)
		if areAll0(matrix) {
			fmt.Printf("First step when all flashed together: %d (Part 2)\n", i+1)
			break
		}
		if i == 99 {
			fmt.Printf("Total flashes after %d steps: %d (Part 1)\n", i+1, flashes)
		}
	}
}

func areAll0(matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 0 {
				return false
			}
		}
	}

	return true
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

type pair struct {
	i int
	j int
}

func step(matrix [][]int) int {
	dx := [8]int{0, 1, 0, -1, 1, 1, -1, -1}
	dy := [8]int{1, 0, -1, 0, 1, -1, 1, -1}

	flashes := 0
	flashed := make(map[pair]bool)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j]++
		}
	}

	stillFlashes := true
	for stillFlashes {
		stillFlashes = false
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				if alreadyFlashed := flashed[pair{i, j}]; !alreadyFlashed && matrix[i][j] > 9 {
					flashed[pair{i, j}] = true
					for k := 0; k < len(dx); k++ {
						if i+dy[k] >= 0 && i+dy[k] < len(matrix) && j+dx[k] >= 0 && j+dx[k] < len(matrix[i]) {
							matrix[i+dy[k]][j+dx[k]]++
							stillFlashes = true
						}
					}
				}
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] > 9 {
				matrix[i][j] = 0
				flashes++
			}
		}
	}

	return flashes
}

func readInput(path string) [][]int {
	input, _ := os.Open(path)
	scanner := bufio.NewScanner(input)

	matrix := make([][]int, 0)

	for scanner.Scan() {
		text := scanner.Text()
		line := make([]int, 0)
		for _, char := range text {
			line = append(line, int(char-'0'))
		}

		matrix = append(matrix, line)
	}

	return matrix
}
