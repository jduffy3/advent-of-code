package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  file, err := os.Open(os.Args[1])
  if err != nil {
    os.Exit(1)
  }
  defer file.Close()

  var forest [][]int
  scanner := bufio.NewScanner(file)
  for start := 0; scanner.Scan(); start++ {
    line := scanner.Text()
    trees := make([]int, len(line))
    for i, t := range strings.Split(line, "") {
      trees[i] = number(t)
    }
    forest = append(forest, trees)
  }

  count, maxScore := solve(forest)
  fmt.Printf("Visible from outside: %d\n", count)
  fmt.Printf("Highest scenic score: %d\n", maxScore)
}

func solve(forest [][]int) (int, int) {
  count, maxScore := 0, 0
  for x := 0; x < len(forest); x++ {
    for y := 0; y < len(forest); y++ {
      visible, score := check(x, y, forest)
      if visible {
        count++
      }
      if score > maxScore {
        maxScore = score
      }
    }
  }
  return count, maxScore
}

func check(x, y int, forest [][]int) (bool, int) {
  tree := forest[x][y]

  left := true
  leftScore := 0
  for start := y - 1; start >= 0; start-- {
    leftScore++
    if forest[x][start] >= tree {
      left = false
      break
    } 
  }

  top := true
  topScore := 0
  for start := x - 1; start >= 0; start-- {
    topScore++
    if forest[start][y] >= tree {
      top = false
      break
    } 
  }

  right := true
  rightScore := 0
  for start := y + 1; start < len(forest); start++ {
    rightScore++
    if forest[x][start] >= tree {
      right = false
      break
    } 
  }

  bottom := true
  bottomScore := 0
  for start := x + 1; start < len(forest); start++ {
    bottomScore++
    if forest[start][y] >= tree {
      bottom = false
      break
    } 
  }

  return (left || top || right || bottom), (leftScore * topScore * rightScore * bottomScore)
}

func number(v string) int {
  n, err := strconv.Atoi(v)
  if err != nil {
    os.Exit(3)
  }
  return n
}

