package main

import (
  "flag"
  "os"
  "os/exec"
  "strings"
  "fmt"
  "io/ioutil"
  "encoding/json"
)

var targetUrl string
var mk = flag.Bool("mk", false, "")
var rm = flag.Bool("rm", false, "")

func main() {
  flag.Parse()
  shortcuts := read()

  if *mk == true {
    // wip
    command := flag.Args()[0]
    url := flag.Args()[1]

    write(shortcuts.Shortcuts, command, url)
  } else if *rm == true {
    // wip
  } else {
    command := flag.Args()[0]
    args := flag.Args()[1:]
    targetShortcut := find(shortcuts.Shortcuts, command)

    arg := strings.Join(args, "-")
    exec.Command("open", targetShortcut.Url+arg).Run()
  }
}

func read() ShortcutsNode {
    jsonFile, err := os.Open("dat.json")
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)
  result := ShortcutsNode{}
  json.Unmarshal([]byte(byteValue), &result)

  return result
}

// wip
func write(currentShortcut []Shortcut, command string, url string) {

  newShortcut := Shortcut {
    Key: command,
    Url: url,
  }

  newContent := append(currentShortcut, newShortcut)
  wrap := ShortcutsNode {
    Shortcuts: newContent,
  }

  file, _ := json.MarshalIndent(wrap, "", " ")
  err := ioutil.WriteFile("dat.json", file, 0644)
  check(err)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func find(shortcuts []Shortcut, target string) Shortcut {
  var targetUrl Shortcut

  for _, v := range shortcuts {
    if v.Key == target {
      targetUrl = v
    }
  }

  return targetUrl
}

func remove(s []int, i int) []int {
    s[len(s)-1], s[i] = s[i], s[len(s)-1]
    return s[:len(s)-1]
}

type ShortcutsNode struct {
  Shortcuts []Shortcut `json:"shortcuts"`
}

type Shortcut struct {
  Key string `json:"key"`
  Url string `json:"url"`
}
