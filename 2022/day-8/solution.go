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

  visibleFromOutside(forest, result)
  fmt.Printf("Visible from outside: %d\n", count(result))
  fmt.Printf("Highest scenic score: %d\n", maxScore(forest))
}

func maxScore(forest [][]int) int {
  max := 0
  for x := range forest {
    for y := range forest {
      s := score(x, y, forest)
      if s > max {
        max = s
      }
    }
  }
  return max
}

func score(row, col int, forest [][]int) int {
  tree := forest[row][col]
  left := 0
  for start := col - 1; start >= 0; start-- {
    left++
    if forest[row][start] >= tree {
      break
    }
  }

  top := 0
  for start := row - 1; start >= 0; start-- {
    top++
    if forest[start][col] >= tree {
      break
    }
  }

  right := 0
  for start := col + 1; start < len(forest); start++ {
    right++
    if forest[row][start] >= tree {
      break
    }
  }

  bottom := 0
  for start := row + 1; start < len(forest); start++ {
    bottom++
    if forest[start][col] >= tree {
      break
    }
  }
  return left * top * right * bottom 
}

func visibleFrom(x, y int, forest [][]int, result [][]bool) {
  if x < 0 || x >= len(forest) {
    return
  }
  if y < 0 || y >= len(forest) {
    return
  }

  tree := forest[x][y]

  left := true
  for start := y - 1; start >= 0; start-- {
    if forest[x][start] >= tree {
      left = false
      break
    } 
  }

  top := true
  for start := x - 1; start >= 0; start-- {
    if forest[start][y] >= tree {
      top = false
      break
    } 
  }

  right := true
  for start := y + 1; start < len(forest); start++ {
    if forest[x][start] >= tree {
      right = false
      break
    } 
  }

  bottom := true
  for start := x + 1; start < len(forest); start++ {
    if forest[start][y] >= tree {
      bottom = false
      break
    } 
  }
  result[x][y] = left || top || right || bottom 
}

func visibleFromOutside(forest [][]int, result [][]bool) {
  for x := 1; x < len(forest); x++ {
    for y := 1; y < len(forest); y++ {
      if !result[x][y] {
        visibleFrom(x, y, forest, result)
      }
    }
  }
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

