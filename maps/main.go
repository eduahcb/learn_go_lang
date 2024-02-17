package main

import "fmt"

func main() {
  fmt.Println("Maps")

  user := map[string]string {
    "name": "Pedro",
    "lastname": "Bosta",
  }

  fmt.Println(user["name"])
}

