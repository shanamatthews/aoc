package main

import (
  "os"
  "bufio"
  "strings"
  "log"
  "fmt"
)

//        them 
//      | A B C
//     -------- 
// m  A | x L W
// e  B | W x L
//    C | L W x

func scoreMatch(p1 string, outcome string) (matchScore int) {
  matchScore = 0
  aScore := 1
  bScore := 2
  cScore := 3

  switch outcome {
    // lose
    case "X":
      matchScore = 0

      switch p1 {
        case "A":
          matchScore = matchScore + cScore
        case "B":
          matchScore = matchScore + aScore
        case "C":
          matchScore = matchScore + bScore
      }

    // draw
    case "Y":
      matchScore = 3

      switch p1 {
        case "A":
          matchScore = matchScore + aScore
        case "B":
          matchScore = matchScore + bScore
        case "C":
          matchScore = matchScore + cScore
      }

    // win
    case "Z":
      matchScore = 6

      switch p1 {
        case "A":
          matchScore = matchScore + bScore
        case "B":
          matchScore = matchScore + cScore
        case "C":
          matchScore = matchScore + aScore
      }
  }

  return
}

func main() {
  file, _ := os.Open("../input.txt")
  defer file.Close()

  var totalScore int
  lineScanner := bufio.NewScanner(file)

  for lineScanner.Scan() {
    words := strings.Fields(lineScanner.Text())
    fmt.Println(words)
    matchScore := scoreMatch(words[0], words[1])
    totalScore = totalScore + matchScore
  }

  if err := lineScanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Println(totalScore)
}
