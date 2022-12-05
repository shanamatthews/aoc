package main

import (
  "os"
  "bufio"
  "log"
  "fmt"
)

const groupSize int = 3

func findCommonItem(group [groupSize]string) rune {
  var itemMaps [groupSize]map[rune]int

  for i, items := range group {
    itemMaps[i] = make(map[rune]int)
    itemMap := itemMaps[i]

    for _, runeValue := range items {
      count := itemMap[runeValue]

      if count == 0 {
        itemMap[runeValue] = 1
      } else {
        itemMap[runeValue] = count + 1
      }
    }
  }
  
  for item := range itemMaps[0] {
    exists := true
    for _, itemMap := range itemMaps {
      if itemMap[item] == 0 {
        exists = false
        break
      }
    }
    if exists {
      return item
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
  var groupItems [groupSize]string

  var totalScore int
  lineScanner := bufio.NewScanner(file)

  for i := 0; lineScanner.Scan(); i++ {
    items := lineScanner.Text()
    groupItems[i % groupSize] = items
    var itemScore int
    
    if i % groupSize == groupSize - 1 {
      commonItem := findCommonItem(groupItems)
      itemScore = scoreItem(commonItem)
    }

    totalScore = totalScore + itemScore
  }

  if err := lineScanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Println(totalScore)
}
