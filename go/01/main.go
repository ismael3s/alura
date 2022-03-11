package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitiringInterval = 2
const delay = 1

func main() {
	for {
		showMenu()

		input := readCommand()

		switch input {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing logs")
			showLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid input")
			os.Exit(-1)
		}
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

	websites := readFromTxt()

	for i := 0; i < monitiringInterval; i++ {
		for _, website := range websites {
			isWebsiteUp(website)
		}

		time.Sleep(time.Second * delay)
	}
	fmt.Println("")

}

func isWebsiteUp(url string) {
	response, _ := http.Get(url)

	if response.StatusCode == 200 {
		fmt.Println("Website", url, "is up!")
		registerLog(url, true)
	} else {
		registerLog(url, false)
		fmt.Println("Website", url, "is down!")
	}
}

func readFromTxt() []string {
	file, err := os.Open("sites.txt")

	if err != nil {
		log.Fatalln("Failed to open websites files", err)
	}

	defer file.Close()

	var sites []string

	reader := bufio.NewReader(file)

	for {
		r, err := reader.ReadString('\n')

		row := strings.TrimSpace(r)

		sites = append(sites, row)

		if err == io.EOF {
			break
		}

	}

	return sites
}

func registerLog(site string, isUp bool) {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Failed to register the log", err)
	}

	defer file.Close()

	message := fmt.Sprintf("%s - %s isUp %t.\n", time.Now().Local().Format("2006-01-02 15:04:05"), site, isUp)

	file.WriteString(message)
}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}

	fmt.Println(string(file))
}
