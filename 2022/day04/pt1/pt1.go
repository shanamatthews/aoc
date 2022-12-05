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

func hasFullOverlap(r1 Range, r2 Range) int {
  fmt.Println(r1)
  fmt.Println(r2)
  return 0
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

    numFullOverlaps = numFullOverlaps + hasFullOverlap(r1, r2)
  }

  if err := lineScanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Println(numFullOverlaps)
}
