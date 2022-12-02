package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type GameResult int
const (
  win = iota
  lose
  draw

  rock = iota + 1
  paper
  scissors
)

func main() {
  file, err := os.Open(os.Args[1])
  if err != nil {
    fmt.Println("error opening file")
    os.Exit(1)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  total_score := 0
  for scanner.Scan() {
    strategy := scanner.Text()
    
    round := strings.Split(strategy, " ")
    total_score += score(round)
  }
  fmt.Println(fmt.Sprintf("total_score %d", total_score))
}

func score(round []string) int {
  shape_values := map[string]int{
    "X": 1,
    "Y": 2,
    "Z": 3,
  }

  my_shape := round[1]
  s := shape_values[my_shape]
  switch result(round) {
    case win: 
      s += 6
    case lose: 
      s += 0
    case draw: 
      s += 3
  }

  return s
}

func result(round []string) GameResult {
  cipher := map[string]int{
    "A": rock,
    "X": rock,
    "B": paper,
    "Y": paper,
    "C": scissors,
    "Z": scissors,
  }

  opponent_shape, my_shape := cipher[round[0]], cipher[round[1]]
  if opponent_shape == my_shape {
    return draw
  }

  switch my_shape { 
  case rock:
    if opponent_shape == scissors {
      return win
    }
  case paper:
    if opponent_shape == rock {
      return win
    }
  case scissors:
    if opponent_shape == paper {
      return win
    }
  }
  return lose
}
