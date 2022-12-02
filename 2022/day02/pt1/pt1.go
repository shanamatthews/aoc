package main

import (
  "os"
  "bufio"
  "strings"
  "log"
  "fmt"
)

//         you
//      | A B C
//     -------- 
// m  A | x L W
// e  B | W x L
//    C | L W x


func scoreMatch(p1 string, p2 string) (matchScore int) {
  matchScore = 0

  switch p2 {
    case "X":
      matchScore = 1

      switch p1 {
        case "A":
          matchScore = matchScore + 3
        case "B":
          matchScore = matchScore + 0
        case "C":
          matchScore = matchScore + 6
      }

    case "Y":
      matchScore = 2

      switch p1 {
        case "A":
          matchScore = matchScore + 6
        case "B":
          matchScore = matchScore + 3
        case "C":
          matchScore = matchScore + 0
      }

    case "Z":
      matchScore = 3

      switch p1 {
        case "A":
          matchScore = matchScore + 0
        case "B":
          matchScore = matchScore + 6
        case "C":
          matchScore = matchScore + 3
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
    // fmt.Println(matchScore)
    totalScore = totalScore + matchScore
  }

  if err := lineScanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Println(totalScore)
}
