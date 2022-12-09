package main

import (
  "os"
  "bufio"
  "strings"
  "unicode"
  "strconv"
  "log"
  "fmt"
)

//      root (FileTree)
//      /  \
//    File FileTree
//          /    \
//        File   File

type FileTree struct {
  Name string
  Size int
  Children map[string]*FileTree
  Parent *FileTree
}

func calculateDirSizes(node *FileTree) int {
  if len((*node).Children) == 0 {
    // base, this is a file
    return (*node).Size
  } else {
    // this is a directory
    // it's size is equal to the sum of all the children
    calculatedSize := 0

    for _, childNode := range (*node).Children {
      calculatedSize = calculatedSize + calculateDirSizes(childNode)
    }

    (*node).Size = calculatedSize
    return calculatedSize
  }
}

func sumSmallDirs(node *FileTree) int {
  if len((*node).Children) == 0 {
    return 0
  } else {
    totalSize := 0

    if (*node).Size < 100000 {
      totalSize = (*node).Size
    }

    for _, childNode := range (*node).Children {
      totalSize = totalSize + sumSmallDirs(childNode)
    }

    return totalSize
  }
}



func main() {
  file, _ := os.Open("../input.txt")
  defer file.Close()

  root := FileTree{}
  root.Name = "/"
  root.Children = make(map[string]*FileTree)
  current := &root
  lineScanner := bufio.NewScanner(file)

  for lineScanner.Scan(){
    line := lineScanner.Text()
    // need to account for cd ..
    if strings.HasPrefix(line, "$ cd") {
      // inialize children for new dir
      // find dir name in currents children
      // change current to that dir
      dirName := strings.Split(line, " ")[2]
      if newDir, contains := (*current).Children[dirName]; contains {
        current = newDir
        (*current).Children = make(map[string]*FileTree)
      } else if dirName == ".." {
        current = current.Parent
      }
    }

    if strings.HasPrefix(line, "dir") {
      dirName := strings.Split(line, " ")[1]
      dir := FileTree{dirName, -1, nil, current}
      
      if _, contains := (*current).Children[dirName]; !contains {
        (*current).Children[dirName] = &dir
      }
    }

    if unicode.IsDigit(rune(line[0])) {
      stuff := strings.Split(line, " ") 
      fileName := stuff[1]
      fileSize, err := strconv.Atoi(stuff[0])
      // add parent to constructor
      file := FileTree{fileName, fileSize, nil, current}
      
      if err != nil {
        log.Fatal(err)
      }

      if _, contains := (*current).Children[fileName]; !contains {
        (*current).Children[fileName] = &file
      }
    }
  }

  // recurse through tree to calculate directory sizes
  calculateDirSizes(&root)

  // recurse through tree again to find directories < 100000
  toRet := sumSmallDirs(&root)


  fmt.Println(root)

  if err := lineScanner.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Println(toRet)
}
