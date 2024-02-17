package main

import "fmt"

func main() {
  // ARITMETICOS
  sum := 1 + 2
  sub := 2 - 1
  division := 10 / 2
  multi := 5 * 5
  restDivision := 10 % 2

  fmt.Println(sum, sub, division, multi, restDivision)
  
  // ATRIBUICAO

  var variable1 string = "hello"
  variable2 := "hello there"
  fmt.Println(variable1, variable2)

  // RELATIONAL

  fmt.Println(1 > 2)
  fmt.Println(1 >= 2)
  fmt.Println(1 == 2)
  fmt.Println(1 <= 2)
  fmt.Println(1 > 2)
  fmt.Println(1 != 2)

  // LOGIC
  truthy, falsy := true, false
  fmt.Println(truthy && falsy)
  fmt.Println(truthy || falsy)
  fmt.Println(!truthy || !falsy)
  
  // UNARY
  number := 10
  number++ // number + 1
  number += 15 // number = number + 15
  fmt.Println(number)
}
