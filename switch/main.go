package main

import "fmt"

func weekDay(number int) string {
  switch number  {
    case 1: 
      return "Domingo"
    case 2: 
      return "Segunda-Feira"
    case 3: 
      return "Terca-Feira"
    case 4: 
      return "Quarta-Feira"
    case 5: 
      return "Quinta-Feira"
    case 6: 
      return "Sexta-Feira"
    case 7: 
      return "Sábado"
    default:
      return "Invalido"
  }
}

func main() {
  day := weekDay(1)
  fmt.Println(day)
}
