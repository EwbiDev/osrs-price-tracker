package main

import (
	"log"

	"github.com/EwbiDev/go-runescape/runescape"
)

func main() {
	client := runescape.NewClient(nil)

	response, err := client.ListGrandExchangeItems("osrs", "a", 1, 1)
	if err != nil {
		log.Fatal("Error:", err)
	}

	log.Println("Total items:", response.Total)
	for _, item := range response.Items {
		log.Println("Name:", item.Name)
		log.Println("Description:", item.Description)
	}
}
