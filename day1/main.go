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
