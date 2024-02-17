package main

import "fmt"

type User struct {
	name string
  age uint8
}

func main() {

  var user User
  user.name = "Carlim"
  user.age = 21
  fmt.Println(user)
  

  secondUser := &User{"Larissa", 22}
  fmt.Println(secondUser)

	newUser := &User{name: "Eduardo", age: 28}
  fmt.Println(newUser)
}
