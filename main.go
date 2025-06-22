package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dracuxan/shodan-cli/shodan"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: shodan-cli <Host IP>")
	}
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
	fmt.Printf("Query Credits: %d\nScan Credits:  %d\n\n", info.QueryCredits, info.ScanCredits)

	hostLookup, err := c.HostInfo(os.Args[1])
	if err != nil {
		log.Fatalf("Error fetching host info: %v", err)
	}

	fmt.Printf("Host Information for %s:\n", os.Args[1])
	for _, i := range hostLookup.Data {
		fmt.Printf("%18s%8d\n", i.IPStr, i.Port)
	}
}
