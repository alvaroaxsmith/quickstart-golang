package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	displayIntroduction()
	for {
		displayMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
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

func startMonitoring() {
	fmt.Println("Monitoring...")
	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was loaded successfully!")
	} else {
		fmt.Println("Site:", site, "is experiencing issues. Status Code:", resp.StatusCode)
	}

}
