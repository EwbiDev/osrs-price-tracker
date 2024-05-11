package main

import (
	"EwbiDev/osrs-price-tracker/pkg/client"
	"fmt"
)

const userAgent = "test - discord@hybrid8513"

func main() {
	geClient := client.NewClient(userAgent)

	responseWiki, err := geClient.GetWikiPrices("24h")
	if err != nil {
		fmt.Println("Error:", err)
	}

	responseOfficial, err := geClient.GetOfficialPrices()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("%v\n", responseWiki.Data["2"])
	fmt.Printf("%v\n", responseOfficial.Data["2"])
	fmt.Printf("%v\n", responseOfficial.JagexTimestamp)
	fmt.Printf("%v\n", responseOfficial.UpdateDetected)
}
