package main

import (
	"log"

	odbclient "github.com/Konzepte-moderner-Softwareentwicklung/ODB/odbClient"
)

func main() {
	client, err := odbclient.NewODBClient("localhost:35000")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer client.Close()

	results := client.GetTelemetry()
	for name, result := range results {
		log.Printf("%s: %+v", name, result)
	}
}
