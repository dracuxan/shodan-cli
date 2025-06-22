package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dracuxan/shodan-cli/shodan"
	"github.com/joho/godotenv"
)

func main() {
	user := os.Getenv("USER")
	path := fmt.Sprintf("/home/%s/.shodan-cli.conf", user)
	if err := godotenv.Load(path); err != nil {
		log.Fatalf("Unable to load env file: %v", err)
	}

	key := os.Getenv("SHODAN_API_KEY")
	c := shodan.New(key)
	info, err := c.APIInfo()
	if err != nil {
		log.Fatalf("Error fetching API info: %v", err)
	}

	fmt.Printf("API Info:\n")
	fmt.Printf("Scan Credits: %d\n", info.ScanCredits)
	fmt.Printf("Query Credits: %d\n", info.QueryCredits)
	fmt.Printf("Monitored IPs: %d\n", info.MonitoredIps)
	fmt.Printf("Plan: %s\n", info.Plan)
	fmt.Printf("HTTPS Enabled: %t\n", info.Https)
	fmt.Printf("Unlocked: %t\n", info.Unlocked)
	fmt.Printf("Unlocked Left: %d\n", info.UnlockedLeft)
	fmt.Printf("Telnet Enabled: %t\n", info.Telnet)
}
