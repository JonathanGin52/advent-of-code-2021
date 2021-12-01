package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  rawData, _ := os.ReadFile("./day1/input.txt")
  data := strings.Split(string(rawData), "\n")

  fmt.Println(part1(data))
  fmt.Println(part2(data))
}

func part1(data []string) int {
  count := 0
  prev, _ := strconv.Atoi(data[0])

  for _, val := range data[1:] {
    curr, _ := strconv.Atoi(val)

    if curr > prev {
      count++
    }

    prev = curr
  }

  return count
}

func part2(data []string) int {
  count := 0
  prevSum := -1

  for i := 1; i < len(data) - 1; i++ {
    prev, _ := strconv.Atoi(data[i - 1])
    curr, _ := strconv.Atoi(data[i])
    next, _ := strconv.Atoi(data[i + 1])

    sum := prev + curr + next

    if sum > prevSum && prevSum != -1 {
      count++
    }

    prevSum = sum
  }

  return count
}
