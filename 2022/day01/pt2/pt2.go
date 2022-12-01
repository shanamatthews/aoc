package main

import (
  "os"
  "bufio"
  "strconv"
  "sort"
  "log"
  "fmt"
)

func main() {
  file, _ := os.Open("../input.txt")
  defer file.Close()

  totals := []int{}
  var currentCalories int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    var line string = scanner.Text()

    if line == "" {
      totals = append(totals, currentCalories)
      currentCalories = 0
    } else {
      lineCals, err := strconv.Atoi(line)

      if err != nil {
        log.Fatal(err)
      }

      currentCalories = currentCalories + lineCals
    }
  }

  totals = append(totals, currentCalories)
  sort.Sort(sort.Reverse(sort.IntSlice(totals)))

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Printf("1: %d 2: %d 3: %d\n", totals[0],totals[1],totals[2])
  fmt.Printf("total: %d\n", totals[0]+totals[1]+totals[2])
}
