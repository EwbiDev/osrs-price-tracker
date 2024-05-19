package controllers

import (
	"EwbiDev/osrs-price-tracker/db"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
)

type ItemController struct {
	queries *db.Queries
	ctx     context.Context
}

func NewItemController(db *db.Queries, ctx context.Context) *ItemController {
	return &ItemController{db, ctx}
}

func (ic *ItemController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idInt, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error id must be int"))
		return
	}

	item, err := ic.queries.SelectItem(ic.ctx, idInt)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("controllers:ItemController:Get - SelectItem " + string(err.Error())))
		return
	}

	itemJson, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("controllers:ItemController:Get - json.Marshal " + string(err.Error())))
		return
	}

	w.Write(itemJson)
}

func (ic *ItemController) List(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	itemParams := populateItemParams(queries)

	items, err := ic.queries.SelectItems(ic.ctx, itemParams)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("controllers:ItemController:List - ListItems " + string(err.Error())))
		return
	}

	itemJson, err := json.Marshal(items)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("controllers:ItemController:List - json.Marshal " + string(err.Error())))
		return
	}

	w.Write(itemJson)
}

func populateItemParams(queries url.Values) db.SelectItemsParams {
	itemParams := db.SelectItemsParams{
		ID:         getParam(queries, "id"),
		Name:       getParam(queries, "name"),
		Icon:       getParam(queries, "icon"),
		TradeLimit: getParam(queries, "trade_limit"),
		Members:    boolishToInt(queries, "members"),
		ItemValue:  getParam(queries, "item_value"),
		LowAlch:    getParam(queries, "low_alch"),
		HighAlch:   getParam(queries, "high_alch"),
		CreatedAt:  getParam(queries, "created_at"),
		UpdatedAt:  getParam(queries, "updated_at"),
	}

	return itemParams
}

func boolishToInt(queries url.Values, param string) *string {
	query := queries.Get(param)

	b, err := strconv.ParseBool(query)
	if err != nil {
		return nil
	}

	var returnVal = "0"
	if b {
		returnVal = "1"
	}

	return &returnVal
}

func getParam(queries url.Values, param string) *string {
	if queries.Has(param) {
		returnString := queries.Get(param)
		return &returnString
	}
	return nil
}
