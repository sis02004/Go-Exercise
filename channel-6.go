package main

import "fmt"

func sum(a, b int) <-chan int {
  out := make(chan int)

  go func(){
    out <- a+b
  }()
  return out
}

func main(){
  c := sum(1,2)
  fmt.Println(<-c)
}
