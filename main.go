package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

func checkHealth(urls []string) {
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("%sERROR: %s - %v%s", colorRed, url, err, colorReset)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("%sERROR: %s returned status %d%s", colorRed, url, resp.StatusCode, colorReset)
			continue
		}

		log.Printf("%sSUCCESS: %s returned status %d%s", colorGreen, url, resp.StatusCode, colorReset)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	urlsEnv := os.Getenv("CHECK_URLS")
	if urlsEnv == "" {
		log.Fatal("No URLs provided. Set CHECK_URLS in the .env file.")
	}

	urls := strings.Split(urlsEnv, ",")

	for {
		log.Println("Starting health check...")
		checkHealth(urls)
		log.Println("Health check completed. Waiting 30 seconds before next check.")
		time.Sleep(30 * time.Second)
	}
}
