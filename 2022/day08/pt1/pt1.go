package main

import (
  "os"
  "bufio"
  "strconv"
  "log"
  "fmt"
)

type MaxDirections struct {
  TallestUp int
  TallestDown int
  TallestLeft int
  TallestRight int
  Height int
  Visible bool
}

func readInputFile(path string) error {
  file, err := os.Open(path)

  if err != nil {
    panic("Cannot open file")
  }

  defer file.Close()

  lineScanner := bufio.NewScanner(file)
  trees := make([][]*MaxDirections, 0)

  // check ->, V
  for i := 0; lineScanner.Scan(); i++ {
    line := lineScanner.Text()
    treesLine := make([]*MaxDirections, 0)

    for j, c := range line {
      treeHeight, err := strconv.Atoi(string(c))

      if err != nil {
        return err
      }

      tree := &MaxDirections { -1, -1, -1, -1, treeHeight, false}

      if i == 0 {
        tree.TallestUp = treeHeight
        tree.Visible = true
      } else {
        prevTallestUp := trees[i-1][j].TallestUp

        if treeHeight > prevTallestUp {
          tree.TallestUp = treeHeight
          tree.Visible = true
        } else {
          tree.TallestUp = prevTallestUp
          tree.Visible = tree.Visible || false 
        }
      }

      if j == 0 {
        tree.TallestLeft = treeHeight
        tree.Visible = true
      } else {
        prevTallestLeft := treesLine[j-1].TallestLeft
        
        if treeHeight > prevTallestLeft {
          tree.TallestLeft = treeHeight
          tree.Visible = true
        } else {
          tree.TallestLeft = prevTallestLeft
          tree.Visible = tree.Visible || false
        }
      }
      
      treesLine = append(treesLine, tree)
    }

    trees = append(trees, treesLine)
  }

  // check <-, ^
  maxTreesI := len(trees) - 1
  for i := maxTreesI; i >= 0; i-- {
    treesLine := trees[i]
    maxTreesJ := len(treesLine) - 1

    for j := maxTreesJ; j >= 0; j-- {
      tree := treesLine[j]
      treeHeight := tree.Height

      if i == maxTreesI {
        tree.TallestDown = treeHeight
        tree.Visible = true
      } else {
        prevTallestDown := trees[i+1][j].TallestDown

        if treeHeight > prevTallestDown {
          tree.TallestDown = treeHeight
          tree.Visible = true
        } else {
          tree.TallestDown = prevTallestDown
          tree.Visible = tree.Visible || false
        }
      }

      if j == maxTreesJ {
        tree.TallestRight = treeHeight
        tree.Visible = true
      } else {
        prevTallestRight := treesLine[j+1].TallestRight
        
        if treeHeight > prevTallestRight {
          tree.TallestRight = treeHeight
          tree.Visible = true
        } else {
          tree.TallestRight = prevTallestRight
          tree.Visible = tree.Visible || false
        }
      }
    }
  }

  // count # visible trees
  numVisible := 0
  for _, treeLine := range trees {
    for _, tree := range treeLine {
      if tree.Visible {
        numVisible = numVisible + 1
//        fmt.Print("x")
      } else {
//        fmt.Print(" ")
      }
    }
//    fmt.Print("\n")
  }

  fmt.Println(numVisible)

  err = lineScanner.Err()
  return err
}

func main() {
  path := "../input.txt"
  err := readInputFile(path)

  if err != nil {
    log.Fatal(err)
  }
}
