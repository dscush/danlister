package main

import (
  "fmt"
  "path/filepath"
  "flag"
  "io/ioutil"
)

func printText(path string, depth int) {
  files, _ := ioutil.ReadDir(path)
  for _, file := range files {
    for i := 0; i < depth; i ++ {
      fmt.Print("  ")
    }
    fmt.Println(file.Name())
    if file.IsDir() {
      printText(filepath.Join(path, file.Name()), depth + 1)
    }
  }
}

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

  fmt.Println(root)
  printText(root, 1)
}

