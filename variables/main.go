package main

import "fmt"

func main() {
  // String
  var myVariable string = "Variable"
  fmt.Println(myVariable)

  myVariable2 := "Variable 2"
  fmt.Println(myVariable2)
 

  var (
    myVariable3 string = "lalala"
    myVariable4 string = "lalala"
  )
  fmt.Println(myVariable3, myVariable4)

  myVariable5, myVmyVariable6 := "Variable 5", "Variable 6"
  fmt.Println(myVariable5, myVmyVariable6)
  
}
