package main
import (
  "fmt"

  "testing"
)

func TestByte_in_dbl_slice (t *testing.T){
  dict := make([][]byte, 2) // создадим dict [[1 2] [3]]
  dict [0] = make([]byte, 2)
  dict [1] = make([]byte, 1)
  dict [0][0] = 1
  dict [0][1] = 2
  dict [1][0] = 3
  fmt.Println("dict  ", dict) //это потом уберем
  char := make([]byte, 2) //char [3]
  char [0] = 1
  char [1] = 2
  fmt.Println("char  ", char)//и это уберем
  //проверяем функцию
  bl, id := byte_in_dbl_slice(dict, char) //в bl и id должны записаться true и 0 соответственно
  fmt.Println("bl ", bl, " id ", id)
  if (bl != true) || (id != 0){ //проверяем так ли это
    t.Error("ОШИБКА, БЛЯТЬ!")
  }
}
