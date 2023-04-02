package main

import (
  "bytes"
  "fmt"
  "strconv"
)

func main() {
  var buffer bytes.Buffer
  for i:=1; i <= 3000; i++ {
    buffer.WriteString("line " + strconv.Itoa(i) + "\n")
  }
  fmt.Println(buffer.String())
}
