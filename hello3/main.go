package main

import "fmt"
import "reflect"

func main() {
  name := "Alvaro"
  age := 29
  version := 0.1
  
  fmt.Println("Hello Mr.",name,", you are", age)
  fmt.Println("This program is in version.", version)
  fmt.Println(reflect.TypeOf(version))
}