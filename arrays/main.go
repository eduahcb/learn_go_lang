package main

import "fmt"


func main() {
  fmt.Println("Arrays and Slices")

  var array1 [5]int

  array1[0] = 1
  fmt.Println(array1)

  array2 := [5]int{1, 2, 3, 4, 5}
  fmt.Println(array2)


  slice := []int{1, 2, 3}
  fmt.Println(slice)

  slice = append(slice, 6)
  fmt.Println(slice)

  slice2 := array2[1:3]
  fmt.Println(slice2)
}
