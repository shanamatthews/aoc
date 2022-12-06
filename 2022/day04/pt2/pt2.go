package main

import (
  "os"
  "bufio"
  "strings"
  "strconv"
  "log"
  "fmt"
)

type Range struct {
  Start int
  End int
}

func hasOverlap(r1 Range, r2 Range) bool {
  if r1.Start < r2.Start {
    if r1.End < r2.Start {
      return false
    } else {
      fmt.Println(r1, r2)
      return true
    }
  } else {
    if r2.End < r1.Start {
      return false
    } else {
      fmt.Println(r1, r2)
      return true
    }
  }
}

func main() {
  file, _ := os.Open("../input.txt")
  defer file.Close()

  var numOverlaps int
  lineScanner := bufio.NewScanner(file)

  for lineScanner.Scan() {
    items := lineScanner.Text()
    ranges := strings.Split(items, ",")

    r1Start, err := strconv.Atoi(strings.Split(ranges[0], "-")[0])
    r1End, err := strconv.Atoi(strings.Split(ranges[0], "-")[1])
    r2Start, err := strconv.Atoi(strings.Split(ranges[1], "-")[0])
    r2End, err := strconv.Atoi(strings.Split(ranges[1], "-")[1])

    if err != nil {
      log.Fatal(err)
    }

    r1 := Range{r1Start, r1End}
    r2 := Range{r2Start, r2End}

    if hasOverlap(r1, r2) {
//      fmt.Println(r1, r2)
      numOverlaps = numOverlaps + 1 
    }
  }

  if err := lineScanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Println(numOverlaps)
}
