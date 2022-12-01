package main

import (
  "fmt"
  "bufio"
  "os"
  "log"
  "strconv"
)

func main() {
  file, _ := os.Open("../input.txt")
  defer file.Close()

  var mostCalories, currentCalories int

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    var line string = scanner.Text()

    if line == "" {
      if currentCalories > mostCalories {
        mostCalories = currentCalories
      }

      currentCalories = 0
    } else {
      lineCals, err := strconv.Atoi(line)

      if err != nil {
        log.Fatal(err)
      }

      currentCalories = currentCalories + lineCals
    }
  }

  if currentCalories > mostCalories {
    mostCalories = currentCalories
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Println(mostCalories)
}
