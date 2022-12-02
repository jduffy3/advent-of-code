package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

const next_elf = ""

func main() {
  file, err := os.Open(os.Args[1])
  if err != nil {
    fmt.Println("error opening file")
    os.Exit(1)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  max_calories, total_calories := 0, 0
  for scanner.Scan() {
    t := scanner.Text()
    switch t {
    case next_elf:
      if total_calories > max_calories {
        max_calories = total_calories
      }
      total_calories = 0
    default:
      i, err := strconv.Atoi(t)
      if err != nil {
        fmt.Println("error converting calories")
        os.Exit(1)
      }
      total_calories += i
    }
  }
  fmt.Println(fmt.Sprintf("max_calories: %d", max_calories))
}
