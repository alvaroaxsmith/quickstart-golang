package main

import "fmt"

func main() {
  name := "Alvaro"
  
  fmt.Println("Hello",name)

  fmt.Print("1 - Start monitoring")
  fmt.Print("2 - Show logs")
  fmt.Print("3 - Exit")

  var comando int

  fmt.Scan(&comando)

  fmt.Println(comando)
}