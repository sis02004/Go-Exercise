package main

import (
  "fmt"
  "runtime"
  "time"
  "sync"
)

func main(){
  runtime.GOMAXPROCS(runtime.NumCPU())

  var data int = 0
  var rwMutex = new(sync.RWMutex)

  go func(){
    for i := 0; i<3; i++{
      rwMutex.Lock()
      data += 1
      fmt.Println("write : ", data)
      time.Sleep(10*time.Millisecond)
      rwMutex.Unlock()
    }
  }()

  go func(){
    for i := 0; i<3;i++{
      rwMutex.RLock()
      fmt.Println("read 1 : ", data)
      time.Sleep(1*time.Second)
      rwMutex.RUnlock()
    }
  }()

  go func(){
    for i := 0; i<3; i++{
      rwMutex.RLock()
      fmt.Println("read 2 : ", data)
      time.Sleep(2*time.Second)
      rwMutex.RUnlock()
    }
  }()

  time.Sleep(10*time.Second)
}
