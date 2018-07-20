package main

import (
  "fmt"
  "strconv"
)

type Person struct{
  name string
  age int
}

func formatString(arg interface{}) string{
  switch arg.(type){
  case Person:
    p := arg.(Person)
    return p.name + " " + strconv.Itoa(p.age)
  case *Person:
    p := arg.(*Person)
    return p.name + " " + strconv.Itoa(p.age)
  default:
    return "Error"
  }
}

func main(){
  fmt.Println(formatString(Person{"Maria", 20}))
  fmt.Println(formatString(&Person{"John", 12}))

  var andrew = new(Person)
  andrew.name = "Andrew"
  andrew.age = 35

  fmt.Println(formatString(andrew))
}
