package main

import "fmt"

func main() {
	name := "Alvaro"

	fmt.Println("Hello", name)

	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")

	var command int

	fmt.Scan(&command)

	fmt.Println("you chose command", command)

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Displaying Logs...")
	case 0:
		fmt.Println("Exiting the program...")
	default:
		fmt.Println("Unknown command")

	}
}
