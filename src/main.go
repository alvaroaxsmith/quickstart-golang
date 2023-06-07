package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitorings = 2
const delay = 5

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
			printLogs()
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
	fmt.Println("")

	return inputCommand
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	sites := readSitesFromFile()

	for i := 0; i < monitorings; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was loaded successfully!")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, "is experiencing issues. Status Code:", resp.StatusCode)
		registerLog(site, false)
	}
}

func readSitesFromFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}

	}

	file.Close()
	return sites
}

func registerLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))

}
