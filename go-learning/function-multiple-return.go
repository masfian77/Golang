package main

import "fmt"

func main() {
  var x1,x2 = show_message(10,5)
  fmt.Println(x1)
  fmt.Println(x2)
}

func show_message(x int, y int)(int,int)  {
  var z = x * y
  var a = x / y
  return z,a
}
