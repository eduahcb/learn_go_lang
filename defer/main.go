package main

import "fmt"

func function1() {
  fmt.Println("exec function 1")
}

func function2() {
  fmt.Println("exec function 2")
}


func main() {
  defer function1()
  function2()
}
