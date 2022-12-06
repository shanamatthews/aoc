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

func hasFullOverlap(r1 Range, r2 Range) bool {
  if r1.Start >= r2.Start && r1.End <= r2.End {
    return true
  } else {
    return false
  }
}

func main() {
  file, _ := os.Open("../input.txt")
  defer file.Close()

  var numFullOverlaps int
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

    if hasFullOverlap(r1, r2) || hasFullOverlap(r2, r1) {
      numFullOverlaps = numFullOverlaps + 1
    }
  }

  if err := lineScanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Println(numFullOverlaps)
}
