package main

import (
  "strings"
  "strconv"
  "bufio"
  "fmt"
  "os"
)

type Crate struct {
  value string
}

func (c Crate) name() string {
  return string(c.value[1])
}

func (c Crate) hasCargo() bool {
  return c.value[0] == '[' &&  c.value[2] == ']'
}

func (c Crate) String() string {
  return c.name()
}

func NewInstruction(s string) Instruction {
  i := strings.Split(s, " ")
  return Instruction {
    move: number(i[1]),
    from: number(i[3]),
    to: number(i[5]),
  }
}

type Instruction struct {
  move int
  from int
  to int
}

func number(a string) int {
  n, err := strconv.Atoi(a)
  if err != nil {
    os.Exit(2)
  }
  return n
}

func main() {
  file, err := os.Open(os.Args[1])
  if err != nil {
    os.Exit(1)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  
  stacks := buildStack(scanner)
  instructions := gatherInstructions(scanner)
  displayTop(move(instructions, stacks, strategy9000))
  displayTop(move(instructions, stacks, strategy9001))
}

func gatherInstructions(scanner *bufio.Scanner) []Instruction {
  var instructions []Instruction
  for scanner.Scan() {
    instructions = append(instructions, NewInstruction(scanner.Text()))
  }
  return instructions
}

func buildStack(scanner *bufio.Scanner) map[int][]Crate {
  stacks := make(map[int][]Crate)
  for scanner.Scan() {
    row := scanner.Text()
    if row == "" {
      break
    }

    stackNum := 1
    for i, j := 0, 3; j <= len(row); {
      crate := Crate { value: row[i:j] }
      if crate.hasCargo() {
        s := stacks[stackNum]
        s = append(s, crate)
        stacks[stackNum] = s
      }
      i = j + 1
      j = i + 3
      stackNum++
    }
  }
  return stacks
}


func strategy9000(from, to []Crate) []Crate {
    for _, c := range from {
      to = append([]Crate { c }, to...)
    }
    return to
}

func strategy9001(from, to []Crate) []Crate {
  return append(append([]Crate{}, from...), to...)
}


func move(instructions []Instruction, stacks map[int][]Crate, strategy func ([]Crate, []Crate)[]Crate) map[int][]Crate {
  s := make(map[int][]Crate)
  for k, v := range stacks {
    s[k] = v
  }
  for _, i := range instructions {
    fromStack, toStack := s[i.from], s[i.to]
    cargo := fromStack[:i.move]

    updatedStack := strategy(cargo, toStack)

    s[i.to], s[i.from] = updatedStack, s[i.from][len(cargo):]
  }
  return s
}


func displayTop(stacks map[int][]Crate) {
  fmt.Print("Top of the stack ")
  for n := 1; stacks[n] != nil; n++ {
    fmt.Printf("%s", stacks[n][0])
  }
  fmt.Println()
}

