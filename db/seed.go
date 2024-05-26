package db

import (
	"EwbiDev/osrs-price-tracker/client"
	"context"
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func populateFromOfficial(ctx context.Context, geClient *client.Client, queries *Queries) error {
	responseOfficial, err := geClient.GetOfficialPrices()
	if err != nil {
		return err
	}

	for _, v := range responseOfficial.Data {
		itemData := InsertItemParams{
			ID:         int64(v.ID),
			Name:       v.Name,
			Icon:       v.Icon,
			TradeLimit: int64(v.Limit),
			Members:    v.Members,
			ItemValue:  int64(v.Value),
			LowAlch:    int64(v.LowAlch),
			HighAlch:   int64(v.HighAlch),
		}

		_, err = queries.InsertItem(ctx, itemData)
		if err != nil {
			return err
		}

		priceData := InsertOfficialPriceParams{
			ItemID:         int64(v.ID),
			Price:          int64(v.Price),
			LastPrice:      int64(v.Last),
			Volume:         int64(v.Volume),
			JagexTimestamp: responseOfficial.JagexTimestamp,
		}

		_, err = queries.InsertOfficialPrice(ctx, priceData)
		if err != nil {
			return err
		}
	}

	getCount, err := queries.CountItems(ctx)
	if err != nil {
		return err
	}

	log.Printf("Generated %v items", getCount)
	return nil
}

func populateFromWiki(ctx context.Context, geClient *client.Client, queries *Queries) error {
	responseWiki, err := geClient.GetWikiPrices("24h")
	if err != nil {
		return err
	}

	for i, v := range responseWiki.Data {
		itemId, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			return err
		}

		wikiData := InsertWikiPriceParams{
			ItemID:          itemId,
			AvgHighPrice:    int64(v.AvgHighPrice),
			HighPriceVolume: int64(v.HighPriceVolume),
			AvgLowPrice:     int64(v.AvgLowPrice),
			LowPriceVolume:  int64(v.LowPriceVolume),
			Timescale:       responseWiki.Period,
		}

		_, err = queries.InsertWikiPrice(ctx, wikiData)
		if err != nil {
			return nil
		}
	}

	return nil
}

func Seed() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
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

	err = populateFromOfficial(ctx, geClient, queries)
	if err != nil {
		log.Fatalf("error seeding official data: %v", err)
	}

	err = populateFromWiki(ctx, geClient, queries)
	if err != nil {
		log.Fatalf("error seeding wiki data: %v", err)
	}
}
