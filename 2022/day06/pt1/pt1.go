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
    var last4 []rune

    for i, c := range line {
      last4 = append(last4, c)

      if len(last4) > 4 {
        last4 = last4[1:]
      }

      if len(last4) == 4 && checkUnique(last4) {
        fmt.Println(i + 1)
        break
      }
    }
  }

  if err := lineScanner.Err(); err != nil {
    log.Fatal(err)
  }
}
