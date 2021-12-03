package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawData, _ := os.ReadFile("./day2/input.txt")
	data := strings.Split(strings.TrimSpace(string(rawData)), "\n")

	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

func part1(data []string) int {
	horizontal := 0
	depth := 0

	for _, command := range data {
		split := strings.Split(command, " ")
		direction := split[0]
		magnitude, _ := strconv.Atoi(split[1])

		if direction == "forward" {
			horizontal += magnitude
		} else if direction == "up" {
			depth -= magnitude
		} else if direction == "down" {
			depth += magnitude
		} else {
			panic("Invalid direction")
		}
	}

	return horizontal * depth
}

func part2(data []string) int {
	horizontal := 0
	depth := 0
	aim := 0

	for _, command := range data {
		split := strings.Split(command, " ")
		direction := split[0]
		magnitude, _ := strconv.Atoi(split[1])

		if direction == "forward" {
			horizontal += magnitude
			depth += aim * magnitude
		} else if direction == "up" {
			aim -= magnitude
		} else if direction == "down" {
			aim += magnitude
		} else {
			panic("Invalid direction")
		}
	}

	return horizontal * depth
}
