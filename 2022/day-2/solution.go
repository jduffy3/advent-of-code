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

  s1, s2 := 0, 0
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    round := strings.Split(scanner.Text(), " ")
    you, me := round[0], round[1]
    s1 += score(shape(you), strategy1(me))
    s2 += score(shape(you), strategy2(you, me))
  }
  fmt.Println(fmt.Sprintf("part one total_score %d", s1))
  fmt.Println(fmt.Sprintf("part two total_score %d", s2))
}

func score(you, me int) int {
  s := value(me)
  switch result(you, me) {
    case win: s += 6
    case lose: s += 0
    case draw: s += 3
  }
  return s
}

func value(shape int) int {
  switch shape {
    case rock: return 1
    case paper: return 2
    case scissors: return 3
    default: return 0
  }
}

func result(you, me int) GameResult {
  if you == me {
    return draw
  }

  switch me { 
  case rock:
    if you == scissors {
      return win
    }
  case paper:
    if you == rock {
      return win
    }
  case scissors:
    if you == paper {
      return win
    }
  }
  return lose
}

func shape(letter string) int {
  switch letter {
    case "A": return rock
    case "B": return paper
    case "C": return scissors
    default: return -1
  }
}

func strategy1(letter string) int {
  switch letter {
    case "X": return rock
    case "Y": return paper
    case "Z": return scissors
    default: return -1
  }
}

func strategy2(you, letter string) int {
  switch letter {
    case "X": return lose_to(you)
    case "Y": return shape(you) //DRAW
    case "Z": return beat(you)
    default: return -1
  }
}

func lose_to(letter string) int {
  switch shape(letter) {
    case rock: return scissors
    case paper: return rock
    case scissors: return paper
    default: return -1
  }
}

func beat(letter string) int {
  switch shape(letter) {
    case rock: return paper
    case paper: return scissors
    case scissors: return rock
    default: return -1
  }
}

