package client

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Client struct {
	BaseURL   string
	UserAgent string
	client    *http.Client
}

type WikiPrices struct {
	Data      map[string]WikiItem `json:"data"`
	Timestamp int                 `json:"timestamp"`
}

type WikiItem struct {
	AvgHighPrice    int `json:"avgHighPrice"`
	HighPriceVolume int `json:"highPriceVolume"`
	AvgLowPrice     int `json:"avgLowPrice"`
	LowPriceVolume  int `json:"lowPriceVolume"`
}

type OfficialItem struct {
	Examine  string `json:"examine"`
	ID       int    `json:"id"`
	Members  bool   `json:"members"`
	LowAlch  int    `json:"lowalch"`
	Limit    int    `json:"limit"`
	Value    int    `json:"value"`
	HighAlch int    `json:"highalch"`
	Icon     string `json:"icon"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Volume   int    `json:"volume"`
	Last     int    `json:"last"`
}

type intermediateOfficialPrices struct {
	UpdateDetected float64                `json:"%UPDATE_DETECTED%"`
	JagexTimestamp int                    `json:"%JAGEX_TIMESTAMP%"`
	Random         map[string]interface{} `json:"-"`
}

type OfficialPrices struct {
	UpdateDetected int
	JagexTimestamp int
	Data           map[string]OfficialItem
}

func NewClient(userAgent string) *Client {
	return &Client{
		UserAgent: userAgent,
		client:    &http.Client{},
	}
}

func (c *Client) get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) GetWikiPrices(period string) (*WikiPrices, error) {
	baseUrl := "https://prices.runescape.wiki/api/v1/osrs/"

	url := baseUrl + period

	body, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var response WikiPrices
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetOfficialPrices() (*OfficialPrices, error) {
	url := "https://chisel.weirdgloop.org/gazproj/gazbot/os_dump.json"

	body, err := c.get(url)
	if err != nil {
		return nil, err
	}

	var interPrices intermediateOfficialPrices
	err = json.Unmarshal(body, &interPrices)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &interPrices.Random)
	if err != nil {
		return nil, err
	}

	updateDetected, jagexTimestamp := interPrices.Random["%UPDATE_DETECTED%"], interPrices.Random["%JAGEX_TIMESTAMP%"]

	delete(interPrices.Random, "%UPDATE_DETECTED%")
	delete(interPrices.Random, "%JAGEX_TIMESTAMP%")

	b, err := json.Marshal(interPrices.Random)
	if err != nil {
		return nil, err
	}

	var prices OfficialPrices
	err = json.Unmarshal(b, &prices.Data)
	if err != nil {
		return nil, err
	}

	v1, ok1 := updateDetected.(float64)
	if !ok1 {
		return nil, errors.New("unable to assert UpdateDetected into float64")
	}
	prices.UpdateDetected = int(v1)

	v2, ok2 := jagexTimestamp.(float64)
	if !ok2 {
		return nil, errors.New("unable to assert JagexTimeStamp into int")
	}
	prices.JagexTimestamp = int(v2)

	return &prices, nil
}
