package main

import (
  "os"
  "bufio"
  "regexp"
  "strings"
  "strconv"
  "log"
  "fmt"
)

const numStacks = 9

func moveItems(source *[]string, destination *[]string, numToMove int) {
  for i := 0; i < numToMove; i++ {
    *destination = append(*destination, (*source)[len(*source) - 1])
    *source = (*source)[:len(*source) - 1]
  }
}

func main() {
  file, _ := os.Open("../input.txt")
  defer file.Close()

  lineScanner := bufio.NewScanner(file)
  var stacks [numStacks][]string
  const tallestStackHeight = 7

  for i := 0; i <= tallestStackHeight && lineScanner.Scan(); i++ {
    re := regexp.MustCompile("    ")
    level := re.ReplaceAllString(lineScanner.Text(), " [-]")
    items := strings.Split(level, " ")

    for i, item := range items {
      re := regexp.MustCompile("\\[(.*)\\]")
      if item == "[-]" {
        continue
      } else {
        stacks[i] = append([]string{re.ReplaceAllString(item, "$1")}, stacks[i]...)
      }
    }
  }

  // burn 2 lines of input
  for i := 0; i <= 1 && lineScanner.Scan(); i++ {
    continue
  }

  if err := lineScanner.Err(); err != nil {
    log.Fatal(err)
  }

  for lineScanner.Scan() {
    instruction := strings.Split(lineScanner.Text(), " ")
    numToMove, err := strconv.Atoi(instruction[1])
    source, err := strconv.Atoi(instruction[3])
    destination, err := strconv.Atoi(instruction[5])

    moveItems(&stacks[source - 1], &stacks[destination - 1], numToMove)

    if err != nil {
      log.Fatal(err)
    }
  }

  for _, stack := range stacks {
    fmt.Println(stack)
  }

  fmt.Println("answer")
}
