package main

import "fmt"
import "os"

func main() {

	displayIntroduction()
	displayMenu()
	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Displaying Logs...")
	case 0:
		fmt.Println("Exiting the program")
		os.Exit(0)
	default:
		fmt.Println("Unknown command")
		os.Exit(-1)
	}
}

func displayIntroduction() {
	name := "Douglas"
	version := 1.2
	fmt.Println("Hello, Mr.", name)
	fmt.Println("This program is in version", version)
}

func displayMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Display Logs")
	fmt.Println("0- Exit Program")
}

func readCommand() int {
	var inputCommand int
	fmt.Scan(&inputCommand)
	fmt.Println("The chosen command was", inputCommand)

	return inputCommand
}
