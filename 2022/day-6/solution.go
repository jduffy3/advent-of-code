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

  startOfPacket := findDistinctIndex(datastream, 4)
  fmt.Printf("start of packet is %d\n", startOfPacket)

  startOfMessage := findDistinctIndex(datastream, 14)
  fmt.Printf("start of message is %d\n", startOfMessage)
}

func findDistinctIndex(datastream string, numDistinct int) int {
  m := make(map[string]int)
  start := 0
  for len(m) < numDistinct {
    letter := string(datastream[start])
    lastUniqueIndex, seenBefore := m[letter]
    if seenBefore {
      m = make(map[string]int)
      start = lastUniqueIndex + 1
      m[string(datastream[start])] = start
    } else {
      m[letter] = start
    }
    start += 1
  }
  return start
}
