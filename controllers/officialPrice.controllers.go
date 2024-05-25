package controllers

import (
	"EwbiDev/osrs-price-tracker/db"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OfficialPriceController struct {
	queries *db.Queries
	ctx     context.Context
}

func NewOfficialPriceController(db *db.Queries, ctx context.Context) *OfficialPriceController {
	return &OfficialPriceController{db, ctx}
}

func (opc *OfficialPriceController) GetByItemId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	itemIdInt, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "error id must be int", http.StatusBadRequest)
		return
	}

	prices, err := opc.queries.SelectOfficialPricesByItem(opc.ctx, itemIdInt)
	if err != nil {
		http.Error(w, "OfficialPriceController.Get - Select: "+err.Error(), http.StatusNotFound)
		return
	}
	if prices == nil {
		prices = []db.OfficialPrice{}
	}

	response := struct {
		Status string             `json:"status"`
		Prices []db.OfficialPrice `json:"prices"`
	}{
		Status: "success",
		Prices: prices,
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "OfficialPriceController.Get - json.Marshal: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(responseJson)
}
