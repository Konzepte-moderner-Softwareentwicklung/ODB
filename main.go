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

	res := client.GetTelemetry()
	log.Printf("Telemetry: %+v", res)
}
