package main

import (
  "fmt"
)

func main() {
  s1 := HelloWorld("hoge")
  fmt.Println(s1)
}

func HelloWorld(s string) string {
  return "hello world, " + s
}