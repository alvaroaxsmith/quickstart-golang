package main

import "fmt"

func main() {
  name := "Alvaro"
  
  fmt.Println("Hello",name)

  fmt.Println("1 - Start monitoring")
  fmt.Println("2 - Show logs")
  fmt.Println("3 - Exit")

  var comando int

  fmt.Scan(&comando)

  fmt.Println("you chose command",comando)
}