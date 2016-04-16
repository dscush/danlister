package main

import (
  "fmt"
  "path/filepath"
  "os"
  "strings"
  "flag"
)

func main() {
  var root, output string
  var help, recursive bool
  flag.StringVar(&root, "path", "", "<path to folder, required>")
  flag.BoolVar(&help, "help", false, "<print help>")
  flag.BoolVar(&recursive, "recursive", false, "<when set, list files recursively.  default is off>")
  flag.StringVar(&output, "output", "text", "<json|yaml|text, default is text>")
  flag.Parse()

  if help {
    flag.PrintDefaults()
    return
  }
  if root == "" {
    fmt.Println("Must provide path.  Try 'danlister --help' for usage")
    return
  }
  
  base_depth := strings.Count(root, "/")

  var printText = func(path string, fi os.FileInfo, _ error) (err error) {
    if path == root {
      fmt.Println(root)
      return
    }
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
  }



  filepath.Walk(root, printText)
}

