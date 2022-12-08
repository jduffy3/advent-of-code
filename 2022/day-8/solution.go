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

	scanner := bufio.NewScanner(file)
  var forest [][]int
  totalTrees := 0
  for start := 0; scanner.Scan(); start++ {
    line := scanner.Text()
    trees := make([]int, len(line))
    for i, t := range strings.Split(line, "") {
      trees[i] = number(t)
    }
    forest = append(forest, trees)
    totalTrees += len(trees)
  }

  result := NewGrid(len(forest))

  maxScore, count := visibleFromOutside(forest, result)
  fmt.Printf("Visible from outside: %d\n", count)
  fmt.Printf("Highest scenic score: %d\n", maxScore)
}

func visibleFrom(x, y int, forest [][]int, result [][]bool) int {
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
  result[x][y] = left || top || right || bottom 
  return leftScore * topScore * rightScore * bottomScore
}

func visibleFromOutside(forest [][]int, result [][]bool) (int, int) {
  maxScore := 0
  for x := 1; x < len(forest); x++ {
    for y := 1; y < len(forest); y++ {
      if !result[x][y] {
        score := visibleFrom(x, y, forest, result)
        if score > maxScore {
          maxScore = score
        }
      }
    }
  }
  return maxScore, count(result)
}

func NewGrid(size int) [][]bool {
  grid := make([][]bool, size)
  for x := range grid {
    grid[x] = make([]bool, size)
    for y := range grid {
      // trees are visible on the outside
      if x == 0 || x == (size - 1) || y == 0 || y == size - 1 {
        grid[x][y] = true
      } else {
        grid[x][y] = false
      }
    }
  }
  return grid
}

func count(result [][]bool) int {
  c := 0
  for x := range result {
    for y := range result {
      if result[x][y] {
        c++
      }
    }
  }
  return c
}

func number(v string) int {
  n, err := strconv.Atoi(v)
  if err != nil {
    os.Exit(3)
  }
  return n
}

