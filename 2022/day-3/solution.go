package main

import (
  "unicode"
  "errors"
  "bufio"
  "fmt"
  "os"
)


func main() {
  file, err := os.Open(os.Args[1])
  if err != nil {
    fmt.Println("error opening file")
    os.Exit(1)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  sum, groupSum := 0, 0
  var group []string
  for scanner.Scan() {
    items := scanner.Text()

    itemType, err := union(items)
    if err != nil {
      os.Exit(1)
    }

    sum += priority(itemType)

    group = append(group, items)
    if len(group) == 3 {
      groupSum += badgePriority(group)
      group = group[3:]
    }
  }
  fmt.Println(fmt.Sprintf("The priority sum is %d", sum))
  fmt.Println(fmt.Sprintf("The group priority sum is %d", groupSum))
}

func union(items string) (rune, error) {
  numItems := len(items)
  compartmentSize := numItems / 2
  c1, c2 := items[0:compartmentSize], items[compartmentSize:]

  m := make(map[rune]bool)
  for _, letter := range c1 {
    m[letter] = true
  }
  for _, letter := range c2 {
    if m[letter] == true {
      return letter, nil
    }
  }

  return 'a', errors.New("None Found")
}

func badgePriority(group []string) int {
  mItemType := make(map[rune]int)
  for _, items := range group {
    set := make(map[rune]bool)
    for _, itemType := range items {
      set[itemType] = true
    }
    for k := range set {
      mItemType[k] = mItemType[k] + 1
      if mItemType[k] == 3 {
        return priority(k)
      }
    }
  }
  return 0
}

func priority(itemType rune) int {
  p := 0
  if unicode.IsUpper(itemType) {
    p += 26
    itemType = unicode.ToLower(itemType)
  }

  p += int(itemType - 'a') + 1
  return p
}

