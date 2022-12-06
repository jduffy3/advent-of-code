package main

import (
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
  scanner.Scan()

  datastream := scanner.Text()
  fmt.Printf("start of packet is %d\n", lastDistinctIndex(datastream, 4))
  fmt.Printf("start of message is %d\n", lastDistinctIndex(datastream, 14))
}

func lastDistinctIndex(datastream string, size int) int {
  index := 0
  m := make(map[string]int)
  for ; len(m) < size; index++ {
    letter := string(datastream[index])

    if lastSeenAt, seenBefore := m[letter]; seenBefore {
      m = make(map[string]int)
      index = lastSeenAt
    } else {
      m[letter] = index
    }
  }
  return index
}

