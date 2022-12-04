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
  
  fullyContainedCount := 0
  for scanner.Scan() {
    sectionPair := strings.Split(scanner.Text(), ",")
    s1, s2 := strings.Split(sectionPair[0], "-"), strings.Split(sectionPair[1], "-")

    if fullyContainsEither(s1, s2) {
      fullyContainedCount += 1
    }
  }
  fmt.Println(fmt.Sprintf("%d assignments fully contains the other", fullyContainedCount))
}

func fullyContainsEither(s1, s2 []string) bool {
  return (start(s1) <= start(s2) && end(s1) >= end(s2)) ||
  (start(s2) <= start(s1) && end(s2) >= end(s1))
}

func start(section []string) int {
  i, err := strconv.Atoi(section[0])
  if err != nil {
    os.Exit(1)
  }
  return i
}

func end(section []string) int {
  i, err := strconv.Atoi(section[1])
  if err != nil {
    os.Exit(1)
  }
  return i
}

