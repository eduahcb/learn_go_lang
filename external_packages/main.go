package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
  err := checkmail.ValidateFormat("test@test.com")

  if err != nil {
    fmt.Println(err)
  }
   
}
