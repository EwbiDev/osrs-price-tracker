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
		queries := setupTestDb(t)
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

	t.Run("returns bad request when invalid id given", func(t *testing.T) {
		queries := setupTestDb(t)
		ctx := context.Background()
		controller := controllers.NewItemController(queries, ctx)

		request, _ := http.NewRequest(http.MethodGet, "items/potato", nil)
		response := httptest.NewRecorder()

		vars := map[string]string{
			"id": "potato",
		}
		request = mux.SetURLVars(request, vars)

		controller.Get(response, request)

		assertEqual(t, response.Code, http.StatusBadRequest)
	})

	t.Run("returns not found when id not found", func(t *testing.T) {
		queries := setupTestDb(t)
		ctx := context.Background()
		controller := controllers.NewItemController(queries, ctx)

		request, _ := http.NewRequest(http.MethodGet, "items/2", nil)
		response := httptest.NewRecorder()

		vars := map[string]string{
			"id": "2",
		}
		request = mux.SetURLVars(request, vars)

		controller.Get(response, request)

		assertEqual(t, response.Code, http.StatusNotFound)
	})
}

func TestListItem(t *testing.T) {
	t.Run("returns item json", func(t *testing.T) {
		queries := setupTestDb(t)
		ctx := context.Background()
		controller := controllers.NewItemController(queries, ctx)

		newItems := []db.InsertItemParams{
			{
				ID:         2,
				Name:       "potato",
				Icon:       "icon/path/potato.jpg",
				TradeLimit: 100,
				Members:    true,
				ItemValue:  200,
				LowAlch:    400,
				HighAlch:   600,
			},
			{
				ID:         3,
				Name:       "tomato",
				Icon:       "icon/path/tomato.jpg",
				TradeLimit: 200,
				Members:    true,
				ItemValue:  111,
				LowAlch:    222,
				HighAlch:   333,
			},
			{
				ID:         4,
				Name:       "lettuce",
				Icon:       "icon/path/lettuce.jpg",
				TradeLimit: 200,
				Members:    false,
				ItemValue:  100,
				LowAlch:    200,
				HighAlch:   300,
			},
		}

		for _, item := range newItems {
			queries.InsertItem(ctx, item)
		}

		testTable := []struct {
			in        string
			resultLen int
		}{
			{"", 3},
			{"?id=2", 1},
			{"?name=potato", 1},
			{"?name=ato", 2},
			{"?icon=steve", 0},
			{"?trade_limit=100", 1},
			{"?trade_limit=200", 2},
			{"?trade_limit=200&members=false", 1},
			{"?trade_limit=200&members=0", 1},
			{"?trade_limit=200&members=0&item_value=100", 1},
			{"?trade_limit=200&members=0&item_value=111", 0},
			{"?members=true", 2},
			{"?members=t", 2},
			{"?members=1", 2},
			{"?members=TRUE", 2},
			{"?members=false", 1},
			{"?members=f", 1},
			{"?members=0", 1},
			{"?members=FALSE", 1},
			{"?members=true&id=3", 1},
			{"?item_value=100", 1},
			{"?item_value=100&low_alch=200", 1},
			{"?item_value=100&low_alch=200&high_alch=300&name=lettuce&trade_limit=200&members=false", 1},
			{"?made_up_value=500", 3},
			{"?not_real_field=100", 3},
		}

		for _, tt := range testTable {
			request, _ := http.NewRequest(http.MethodGet, tt.in, nil)
			response := httptest.NewRecorder()

			controller.List(response, request)

			var got []db.Item
			json.Unmarshal(response.Body.Bytes(), &got)

			assertEqual(t, response.Code, http.StatusOK)
			assertEqual(t, len(got), tt.resultLen)
		}
	})
}

func assertEqual(t *testing.T, got any, want any) {
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func setupTestDb(t *testing.T) *db.Queries {
	dbInit, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Errorf("error opening database: %v", err)
	}
	schema, err := os.ReadFile("../db/schema.sql")
	if err != nil {
		t.Errorf("failed to read schema.sql: %v", err)
	}
	dbInit.Exec(string(schema))

	return db.New(dbInit)
}
