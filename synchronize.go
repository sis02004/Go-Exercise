package main

import(
  "fmt"
  "runtime"
  "time"
)

func main(){
  runtime.GOMAXPROCS(runtime.NumCPU())

  var data = []int{}

  go func(){
    for i := 0; i<1000; i++{
      data = append(data,1)

      runtime.Gosched()
    }
  }()

  go func(){
    for i :=0;i<1000;i++{
      data = append(data, 1)

      runtime.Gosched()
    }
  }()

  time.Sleep(2*time.Second)

  fmt.Println(len(data))
}
