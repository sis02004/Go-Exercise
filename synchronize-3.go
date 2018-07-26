package main

import(
  "fmt"
  "runtime"
  "sync"
  "time"
)

func main(){
  runtime.GOMAXPROCS(runtime.NumCPU())

  var data = []int{}
  var mutex = new(sync.Mutex)

  go func(){
    for i:=0;i<1000;i++{
      mutex.Lock()
      data = append(data, 1)
      mutex.Unlock()

      runtime.Gosched()
    }
  }()

  go func(){
    for i := 0; i<1000;i++{
      mutex.Lock()
      data = append(data, 1)
      mutex.Unlock()

      runtime.Gosched()
    }
  }()

  time.Sleep(2*time.Second)

  fmt.Println(len(data))
}
