package main

import (
  "os"
  "bufio"
  "log"
  "fmt"
)

func findCommonItem(items string) rune {
  var firstHalf map[rune]int
  var secondHalf map[rune]int
  firstHalf = make(map[rune]int)
  secondHalf = make(map[rune]int)

  for i, runeValue := range items {
    if i<= len(items)/2 - 1 {
      count := firstHalf[runeValue]
      if count == 0 {
        firstHalf[runeValue] = 1
      } else {
        firstHalf[runeValue] = count + 1
      }
    } else {
      count := secondHalf[runeValue]
      if count == 0 {
        secondHalf[runeValue] = 1
      } else {
        secondHalf[runeValue] = count + 1
      }
    }
  }

  for key := range firstHalf {
    if secondHalf[key] != 0 {
      return key
    }
  }
   
  return 0
}

func scoreItem(item rune) int {
  // 97 -> 1
  // 65 -> 27
  if item > 95 {
    return int(item - 96)
  } else {
    return int(item - 38)
  }
}

func main() {
  file, _ := os.Open("../input.txt")
  defer file.Close()

  var totalScore int
  lineScanner := bufio.NewScanner(file)

  for lineScanner.Scan() {
    items := lineScanner.Text()
//    fmt.Println(items)
    commonItem := findCommonItem(items)
    itemScore := scoreItem(commonItem)
//    fmt.Println(itemScore)
    totalScore = totalScore + itemScore
  }

  if err := lineScanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Println(totalScore)
}
