package main

import(
  "fmt"
  "time"
)

func main(){
  c1 := make(chan int)

  go func(){
    for{
      i := <-c1
      fmt.Println("c1 : ", i)
      time.Sleep(100*time.Millisecond)
    }
  }()
  go func(){
    for{
      c1 <- 20
      time.Sleep(500*time.Millisecond)
    }
  }()

  go func(){
    for{
      select{
        case c1<-10:
        case s:= <-c1:
          fmt.Println("c1 : ", s)
      }
    }
  }()

  time.Sleep(10*time.Second)
}
