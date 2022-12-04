package main

import (
  "strings"
  "strconv"
  "bufio"
  "fmt"
  "os"
)


func main() {
  file, err := os.Open(os.Args[1])
  if err != nil {
    os.Exit(1)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  
  containsCount, overlapCount := 0, 0
  for scanner.Scan() {
    sectionPair := strings.Split(scanner.Text(), ",")
    s1, s2 := strings.Split(sectionPair[0], "-"), strings.Split(sectionPair[1], "-")

    if fullyContains(s1, s2) || fullyContains(s2, s1) {
      containsCount += 1
    }

    if containsOverlap(s1, s2) {
      overlapCount += 1
    }
  }
  fmt.Println(fmt.Sprintf("Contains count: %d, Overlap count :%d", containsCount, overlapCount))
}

func containsOverlap(s1, s2 []string) bool {
  return start(s1) <= end(s2) && end(s1) >= start(s2)
}

func fullyContains(s1, s2 []string) bool {
  return (start(s1) <= start(s2) && end(s1) >= end(s2))
}

func start(section []string) int {
  return parseInt(section[0])
}

func end(section []string) int {
  return parseInt(section[1])
}

func parseInt(s string) int {
  i, err := strconv.Atoi(s)
  if err != nil {
    os.Exit(1)
  }
  return i
}

