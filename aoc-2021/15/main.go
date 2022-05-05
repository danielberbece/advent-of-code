package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	costMatrix := readInput("./input")
	findLowestRiskPathCost(costMatrix, 1)
	newCostMatrix := enlargeMatrix(costMatrix)
	findLowestRiskPathCost(newCostMatrix, 2)
}

type position struct {
	x int
	y int
}

type Item struct {
	value    position
	priority int
	index    int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func findLowestRiskPathCost(costMatrix [][]int, part int) {
	destY := len(costMatrix) - 1
	destX := len(costMatrix[0]) - 1

	dx := [4]int{1, -1, 0, 0}
	dy := [4]int{0, 0, 1, -1}

	seenMatrix := createMatrix(len(costMatrix), len(costMatrix[0]), 0)

	pq := make(PriorityQueue, 0)
	pq = append(pq, &Item{
		value:    position{x: 0, y: 0},
		priority: 0,
		index:    0,
	})

	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if item.value.x == destX && item.value.y == destY {
			fmt.Printf("(Part %d) Cost to get to destination is: %d\n", part, item.priority)
			return
		}
		if seenMatrix[item.value.y][item.value.x] == 0 {
			seenMatrix[item.value.y][item.value.x] = 1

			for k := 0; k < 4; k++ {
				neighbourY := item.value.y + dy[k]
				neighbourX := item.value.x + dx[k]

				if neighbourY >= 0 && neighbourY <= destY && neighbourX >= 0 && neighbourX <= destX {
					newItem := &Item{
						value:    position{x: neighbourX, y: neighbourY},
						priority: item.priority + costMatrix[neighbourY][neighbourX],
					}
					heap.Push(&pq, newItem)
				}
			}
		}
	}
}

func enlargeMatrix(matrix [][]int) [][]int {
	newMatrix := make([][]int, 5*len(matrix))

	for i := 0; i < len(newMatrix); i++ {
		newMatrix[i] = make([]int, 5*len(matrix[0]))
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < len(matrix); k++ {
				for l := 0; l < len(matrix[k]); l++ {
					value := (matrix[k][l] + i + j)
					if value > 9 {
						value %= 9
					}
					newMatrix[i*len(matrix)+k][j*len(matrix[0])+l] = value
				}
			}
		}
	}

	return newMatrix
}

func createMatrix(y int, x int, initialValue int) [][]int {
	m := make([][]int, y)

	for i := 0; i < y; i++ {
		m[i] = make([]int, x)
		for j := 0; j < x; j++ {
			m[i][j] = initialValue
		}
	}

	return m
}

func readInput(path string) [][]int {
	input, _ := os.Open(path)
	scanner := bufio.NewScanner(input)

	m := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		mapLine := make([]int, 0)
		for _, n := range line {
			mapLine = append(mapLine, int(n-'0'))
		}
		m = append(m, mapLine)
	}

	return m
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
