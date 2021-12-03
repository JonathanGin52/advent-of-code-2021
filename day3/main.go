package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawData, _ := os.ReadFile("./day3/input.txt")
	data := strings.Split(strings.TrimSpace(string(rawData)), "\n")

	fmt.Println(part1(data))
}

func part1(data []string) int {
	gammaStr := ""
	epsilonStr := ""
	onesCount, zeroesCount := tally(data)

	for i := range onesCount {
		if onesCount[i] > zeroesCount[i] {
			gammaStr += "1"
			epsilonStr += "0"
		} else {
			gammaStr += "0"
			epsilonStr += "1"
		}
	}

	gamma, _ := strconv.ParseInt(gammaStr, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonStr, 2, 64)

	return int(gamma * epsilon)
}

func tally(data []string) ([12]int, [12]int) {
	var onesCount [12]int
	var zeroesCount [12]int

	for _, entry := range data {
		for i, bit := range entry {
			if bit == '1' {
				onesCount[i]++
			} else {
				zeroesCount[i]++
			}
		}
	}

	return onesCount, zeroesCount
}
