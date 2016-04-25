package main

import (
  "os"
  "fmt"
  "path/filepath"
  "flag"
  "io/ioutil"
  "time"
  "encoding/json"
)

type FileInfo struct {
  Name string
  Size int64
  ModTime time.Time
  IsDir bool
  IsLink bool
  LinksTo string
  Children []FileInfo
}

func makeFileInfo(file os.FileInfo, path string, recursive bool) FileInfo {
  children := make([]FileInfo, 0)
  if recursive && file.IsDir(){
    children = getChildren(path, recursive)
  }
  isLink := file.Mode()&os.ModeSymlink == os.ModeSymlink
  linksTo := ""
  if isLink {
    linksTo, _ = os.Readlink(path)
  }
  return FileInfo{
    file.Name(),
    file.Size(),
    file.ModTime(),
    file.IsDir(),
    isLink,
    linksTo,
    children,
  }
}

func getChildren(path string, recursive bool) []FileInfo {
  var children []FileInfo
  childrenFileInfo, _ := ioutil.ReadDir(path)
  for _, f := range childrenFileInfo {
    children = append(children, makeFileInfo(f, filepath.Join(path, f.Name()), recursive))
  }
  return children
}

func printTabs(depth int) {
  for i := 0; i < depth; i ++ {
    fmt.Print("  ")
  }
}
func printText(path string, depth int, recursive bool) {
  files, _ := ioutil.ReadDir(path)
  for _, file := range files {
    printTabs(depth)
    if file.Mode()&os.ModeSymlink == os.ModeSymlink {
      linksTo, _ := os.Readlink(filepath.Join(path, file.Name()))
      fmt.Printf("%v* (%v)\n", file.Name(), linksTo)
    } else {
      fmt.Println(file.Name())
    }
    if recursive && file.IsDir() {
      printText(filepath.Join(path, file.Name()), depth + 1, recursive)
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

  if output == "json" || output == "yaml" {
    fileTree := getChildren(root, recursive)
    if output == "json" {
      f, _ := json.MarshalIndent(fileTree, "", "  ")
      fmt.Println(string(f))
    } else {
      fmt.Println("TODO: yaml")
    }
  } else {
    fmt.Println(root)
    printText(root, 1, recursive)
  }
}

