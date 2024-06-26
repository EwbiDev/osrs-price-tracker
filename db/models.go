// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"
)

type Item struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Icon       string    `json:"icon"`
	TradeLimit int64     `json:"trade_limit"`
	Members    bool      `json:"members"`
	ItemValue  int64     `json:"item_value"`
	LowAlch    int64     `json:"low_alch"`
	HighAlch   int64     `json:"high_alch"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OfficialPrice struct {
	ID             int64     `json:"id"`
	ItemID         int64     `json:"item_id"`
	Price          int64     `json:"price"`
	LastPrice      int64     `json:"last_price"`
	Volume         int64     `json:"volume"`
	JagexTimestamp time.Time `json:"jagex_timestamp"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type WikiPrice struct {
	ID              int64     `json:"id"`
	ItemID          int64     `json:"item_id"`
	AvgHighPrice    int64     `json:"avg_high_price"`
	HighPriceVolume int64     `json:"high_price_volume"`
	AvgLowPrice     int64     `json:"avg_low_price"`
	LowPriceVolume  int64     `json:"low_price_volume"`
	Timescale       string    `json:"timescale"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
