package main

import (
  "fmt"
  "path/filepath"
  "os"
  "strings"
)

func main() {
  root, _ := os.Getwd()
  fmt.Println(root)
  base_depth := strings.Count(root, "/")
  filepath.Walk(root, func(path string, fi os.FileInfo, _ error) (err error) {
    depth := strings.Count(path, "/") - base_depth
    for i := 0; i < depth; i ++ {
      fmt.Print("  ")
    }
    name := fi.Name()
    if fi.IsDir() {
      name += "/"
    }
    fmt.Println(name)
    return
  })
}
