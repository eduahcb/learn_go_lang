package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func multipleReturns(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

func main() {
	result := sum(2, 5)
	fmt.Println(result)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Hello There!")

  sum, sub := multipleReturns(2, 7)
  fmt.Println(sum, sub)
}
