package main

import(
  "fmt"
  "runtime"
)

func main(){
  runtime.GOMAXPROCS(1)

  done := make(chan bool, 2)
  count := 4

  go func(){
    for i := 0; i < count; i++{
      done <- true
      fmt.Println("고루틴 : ", i)
    }
  }()

  for i := 0; i < count; i++{
    <-done
    fmt.Println("메인 함수 : ", i)
  }
}
