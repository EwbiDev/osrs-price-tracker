package controllers_test

import (
	"EwbiDev/osrs-price-tracker/controllers"
	"EwbiDev/osrs-price-tracker/db"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func TestGETItem(t *testing.T) {
	t.Run("returns item json", func(t *testing.T) {
		dbInit, err := sql.Open("sqlite3", ":memory:")
		if err != nil {
			t.Errorf("error opening database: %v", err)
		}
		schema, err := os.ReadFile("../db/schema.sql")
		if err != nil {
			t.Errorf("failed to read schema.sql: %v", err)
		}
		dbInit.Exec(string(schema))
		queries := db.New(dbInit)

		ctx := context.Background()
		controller := controllers.NewItemController(queries, ctx)

		request, _ := http.NewRequest(http.MethodGet, "items/2", nil)
		response := httptest.NewRecorder()

		newItem := db.InsertItemParams{
			ID:         2,
			Name:       "potato",
			Icon:       "icon/path",
			TradeLimit: 777,
			Members:    true,
			ItemValue:  111,
			LowAlch:    222,
			HighAlch:   333,
		}
		queries.InsertItem(ctx, newItem)

		vars := map[string]string{
			"id": "2",
		}
		request = mux.SetURLVars(request, vars)

		controller.Get(response, request)

		var got db.Item
		json.Unmarshal(response.Body.Bytes(), &got)
		want := &newItem


		assertEqual(t, response.Code, http.StatusOK)
		assertEqual(t, got.ID, want.ID)
		assertEqual(t, got.Name, want.Name)
		assertEqual(t, got.Icon, want.Icon)
		assertEqual(t, got.TradeLimit, want.TradeLimit)
		assertEqual(t, got.Members, want.Members)
		assertEqual(t, got.ItemValue, want.ItemValue)
		assertEqual(t, got.LowAlch, want.LowAlch)
		assertEqual(t, got.HighAlch, want.HighAlch)
	})
}

func assertEqual(t *testing.T, got any, want any) {
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
