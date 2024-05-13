package controllers

import (
	"EwbiDev/osrs-price-tracker/db"
	"context"
	"encoding/json"
	"net/http"
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

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error id must be int"))
		return
	}

	item, err := ic.queries.SelectItem(ic.ctx, id)
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
