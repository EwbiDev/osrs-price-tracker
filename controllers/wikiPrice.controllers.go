package controllers

import (
	"EwbiDev/osrs-price-tracker/db"

	"context"
)

type WikiPriceController struct {
	queries *db.Queries
	ctx     context.Context
}

func NewWikiPriceController(db *db.Queries, ctx context.Context) *OfficialPriceController {
	return &OfficialPriceController{db, ctx}
}

