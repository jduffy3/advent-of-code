package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "sort"
)

const nextElf = ""

func main() {
  file, err := os.Open(os.Args[1])
  if err != nil {
    fmt.Println("error opening file")
    os.Exit(1)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  var calories []int
  totalCalories := 0
  for scanner.Scan() {
    line := scanner.Text()
    switch line {
    case nextElf:
      calories = append(calories, totalCalories)
      totalCalories = 0
    default:
      i, err := strconv.Atoi(line)
      if err != nil {
        os.Exit(1)
      }
      totalCalories += i
    }
  }
  sort.Sort(sort.Reverse(sort.IntSlice(calories)))
  top3Total := 0
  for _, v := range calories[:3] {
    top3Total += v
  }
  fmt.Println(fmt.Sprintf("Top 3 total: %d", top3Total))
}
