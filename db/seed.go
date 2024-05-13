package db

import (
	"EwbiDev/osrs-price-tracker/client"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func populateItems(ctx context.Context, geClient *client.Client, queries *Queries) error {
	responseOfficial, err := geClient.GetOfficialPrices()
	if err != nil {
		return err
	}

	for _, v := range responseOfficial.Data {
		insertData := InsertItemParams{
			ID:         int64(v.ID),
			Name:       v.Name,
			Icon:       v.Icon,
			TradeLimit: int64(v.Limit),
			Members:    v.Members,
			ItemValue:  int64(v.Value),
			LowAlch:    int64(v.LowAlch),
			HighAlch:   int64(v.HighAlch),
		}

		_, err = queries.InsertItem(ctx, insertData)
		if err != nil {
			return err
		}

		getCount, err := queries.CountItems(ctx)
		if err != nil {
			return err
		}

		fmt.Printf("Generated %v items", getCount)
	}
	return nil
}

func Seed() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	userAgent := os.Getenv("USER_AGENT")
	geClient := client.NewClient(userAgent)
	if err != nil {
		log.Fatalf("error initiating database: %v", err)
	}

	ctx := context.Background()
	dbInit, err := sql.Open("sqlite3", "db/db.db")
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	queries := New(dbInit)

	populateItems(ctx, geClient, queries)
}
