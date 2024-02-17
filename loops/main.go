package main

import (
	"fmt"
)


func main() {
  
  // i := 0

  // for i < 10 {
  //   i++
  //   fmt.Println("Incriment i")
  //   time.Sleep(time.Second)
  // }
  //
  // fmt.Println(i)


  for j := 0; j < 10; j++ {
    fmt.Println("Incrment j", j)
  }

  names := [3]string{"Carlim", "Fernando", "Zezé"}

  for _, name := range names {
    fmt.Println(name)
  }
}
