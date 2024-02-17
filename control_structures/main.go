package main

import "fmt"

func main() {
  fmt.Println("Control Structures") 

  number := 10

  if number > 15 {
    fmt.Println("Maior que 15")
  } else {
    fmt.Println("Menor ou igual a 15") 
  }

  if anotherNumber := number; anotherNumber > 0 {
    fmt.Println("Número é maior que zero")
  }
}
