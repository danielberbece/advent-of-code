package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	lanternfish := getLanternfish("./input")
	simulation(lanternfish, 80)
	fmt.Println("Part 1:", countFish(lanternfish))
	simulation(lanternfish, 256-80)
	fmt.Println("Part 2:", countFish(lanternfish))
}

func simulation(lanternfish []int, days int) {
	for i := 0; i < days; i++ {
		births := lanternfish[0]
		for j := 0; j < len(lanternfish)-1; j++ {
			lanternfish[j] = lanternfish[j+1]
		}

		lanternfish[6] += births
		lanternfish[8] = births
	}
}

func getLanternfish(path string) []int {
	input, _ := ioutil.ReadFile(path)
	timers := strings.Split(string(input), ",")

	lanternfish := make([]int, 9)

	for _, timer := range timers {
		number, _ := strconv.Atoi(timer)
		lanternfish[number]++
	}

	return lanternfish
}

func countFish(lanternfish []int) int {
	count := 0

	for _, fish := range lanternfish {
		count += fish
	}

	return count
}
