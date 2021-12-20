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
	fmt.Println(part2(data))
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

func part2(data []string) int {
	oxygenCandidates, co2Candidates := data, data

	for currentColumn := 0; currentColumn < 12; currentColumn++ {
		oxygenCandidates = segment(oxygenCandidates, currentColumn, '1')
		if len(oxygenCandidates) == 1 {
			fmt.Println(oxygenCandidates)
			break
		}
	}
	for currentColumn := 0; currentColumn < 12; currentColumn++ {
		co2Candidates = segment(co2Candidates, currentColumn, '0')
		if len(co2Candidates) == 1 {
			fmt.Println(co2Candidates)
			break
		}
	}

	oxygenGeneratorRating, _ := strconv.ParseInt(oxygenCandidates[0], 2, 32)
	co2ScrubberRating, _ := strconv.ParseInt(co2Candidates[0], 2, 32)

	return int(oxygenGeneratorRating * co2ScrubberRating)
}

func segment(input []string, column int, bitCriteria byte) (output []string) {
	oneCount, zeroCount := tallyColumn(input, column)
	var targetValue byte

	if oneCount > zeroCount {
		if bitCriteria == '1' {
			targetValue = '1'
		} else {
			targetValue = '0'
		}
	} else if oneCount < zeroCount {
		if bitCriteria == '1' {
			targetValue = '0'
		} else {
			targetValue = '1'
		}
	} else {
		targetValue = bitCriteria
	}


	for _, entry := range input {
		if entry[column] == targetValue {
			output = append(output, entry)
		}
	}

	return
}


func tallyColumn(data []string, column int) (oneCount, zeroCount int) {
	for _, entry := range data {
		if entry[column] == '1' {
			oneCount++
		} else {
			zeroCount++
		}
	}

	return
}
