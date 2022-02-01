package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	caveMap := readCaveMap("./input")
	n := generatePaths(caveMap, 1)
	fmt.Println("Total paths part 1:", n)
	n = generatePaths(caveMap, 2)
	fmt.Println("Total paths part 2:", n)
}

func generatePaths(graph map[string][]string, part int) int {
	seen := make(map[string]int)
	for cave := range graph {
		seen[cave] = 0
	}
	seen["start"] = 2

	return generatePathsAux(graph, seen, "start", part)
}

func generatePathsAux(graph map[string][]string, seen map[string]int, currentCave string, part int) int {
	if currentCave == "end" {
		return 1
	} else {
		n := 0
		for _, neighbourCave := range graph[currentCave] {
			if canVisit(neighbourCave, seen, part) {
				seen[neighbourCave]++
				n += generatePathsAux(graph, seen, neighbourCave, part)
				seen[neighbourCave]--
			}
		}

		return n
	}
}

func canVisit(cave string, seen map[string]int, part int) bool {
	if isSmallCave(cave) && seen[cave] > 0 {
		if part == 1 {
			return false
		} else if seen[cave] == 1 && noDoubleVisitsSmallCaves(seen) {
			return true
		}

		return false
	}

	return true
}

func noDoubleVisitsSmallCaves(seen map[string]int) bool {
	for cave := range seen {
		if cave != "start" && isSmallCave(cave) && seen[cave] > 1 {
			return false
		}
	}

	return true
}

func isSmallCave(cave string) bool {
	return cave[0] >= 'a' && cave[0] <= 'z'
}

func readCaveMap(path string) map[string][]string {
	input, _ := os.Open(path)
	scanner := bufio.NewScanner(input)

	caveMap := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		caves := strings.Split(line, "-")

		if neighbours, exists := caveMap[caves[0]]; exists {
			neighbours = append(neighbours, caves[1])
			caveMap[caves[0]] = neighbours
		} else {
			neighbours := make([]string, 0)
			neighbours = append(neighbours, caves[1])
			caveMap[caves[0]] = neighbours
		}

		if neighbours, exists := caveMap[caves[1]]; exists {
			neighbours = append(neighbours, caves[0])
			caveMap[caves[1]] = neighbours
		} else {
			neighbours := make([]string, 0)
			neighbours = append(neighbours, caves[0])
			caveMap[caves[1]] = neighbours
		}
	}

	return caveMap
}
