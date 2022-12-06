package main

import (
  "os"
  "bufio"
  "log"
  "fmt"
)

func checkUnique(chars []rune) bool {
  var dict map[rune]int
  dict = make(map[rune]int)

  fmt.Println(string(chars))

  for _, c := range chars {
    if dict[c] == 0 {
      dict[c] = 1
    } else {
      return false
    }
  }

  return true
}

func main() {
  file, _ := os.Open("../input.txt")
  defer file.Close()

  lineScanner := bufio.NewScanner(file)

  for lineScanner.Scan(){
    line := lineScanner.Text()
    var last14 []rune

    for i, c := range line {
      last14 = append(last14, c)

      if len(last14) > 14 {
        last14 = last14[1:]
      }

      if len(last14) == 14 && checkUnique(last14) {
        fmt.Println(i + 1)
        break
      }
    }
  }

  if err := lineScanner.Err(); err != nil {
    log.Fatal(err)
  }
}
