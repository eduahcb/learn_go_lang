package main

import (
	"errors"
	"fmt"
)

func main() {
  
  // INT

  // int8, intint16, intint32, intint64
  
  // int: Este tipo de inteiro tem um tamanho específico dependente da arquitetura do sistema onde o código está sendo executado. Em sistemas de 32 bits, int tem 32 bits de tamanho, enquanto em sistemas de 64 bits, int tem 64 bits de tamanho
  var number int = 10000
  fmt.Println(number)


  // int8: É um inteiro de 8 bits com sinal. O intervalo de valores possíveis é de -128 a 127
  var number2 int8 = 10
  fmt.Println(number2)

  // int16: É um inteiro de 16 bits com sinal. O intervalo de valores possíveis é de -32768 a 32767.
  var number3 int16 = 10000
  fmt.Println(number3)


  // int32: É um inteiro de 32 bits com sinal. O intervalo de valores possíveis é de -2147483648 a 2147483647
  var number4 int32 = 100000
  fmt.Println(number4)

  // int64: É um inteiro de 64 bits com sinal. O intervalo de valores possíveis é de -9223372036854775808 a 9223372036854775807

  var number5 int64 = 100000
  fmt.Println(number5)

  // REAL Numbers
  // float32, flfloat64

  var myRealNumber float32 = 123.45
  fmt.Println(myRealNumber)


  myRealNumber2 := 1212.21
  fmt.Println(myRealNumber2)


  // STRINGS

  var str string = "text"
  fmt.Println(str)


  str2 := "kkkk"
  fmt.Println(str2)

  char := 'B'
  fmt.Println(char)


  // ERRORS
  var erro error = errors.New("my error")
  fmt.Println(erro)
}
