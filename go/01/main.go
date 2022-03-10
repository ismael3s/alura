package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	showMenu()

	input := readCommand()

	switch input {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Showing logs")
	case 0:
		fmt.Println("Exiting...")
	default:
		fmt.Println("Invalid input")
	}

}

func showMenu() {
	fmt.Println("1. Start monitoring")
	fmt.Println("2. Show logs")
	fmt.Println("0. Exit")
}

func readCommand() int {
	var input int
	fmt.Scanln(&input)

	return input
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	website := "https://www.google.com"

	response, err := http.Get(website)

	if err != nil {
		log.Fatal("Website is down")
	}

	fmt.Println(response.Status)
}
