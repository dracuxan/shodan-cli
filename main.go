package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Unable to load env file: %v", err)
	}

	key := os.Getenv("SHODAN_API_KEY")
	resp, err := http.Get("https://api.shodan.io/shodan/ports?key=" + key)
	if err != nil {
		log.Fatalf("Error fetching data from Shodan: %v", err)
	}

	defer resp.Body.Close()

	var ports []int

	if err := json.NewDecoder(resp.Body).Decode(&ports); err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
	}
	log.Println("Ports found on Shodan:", ports[:10])
}
