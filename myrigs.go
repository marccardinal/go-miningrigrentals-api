package miningrigrentals

import (
	"encoding/json"
)

type MyRigs struct {
	Id           uint64 `json:"id,string"`
	Name         string `json:"name"`
	Rpi          string
	Type         string
	Online       uint8 `json:"online,string"`
	Price        float64 `json:"price,string"`
	PriceHour    float64 `json:"price_hr,string"`
	MinHours     uint16 `json:"minhrs,string"`
	MaxHours     uint16 `json:"maxhrs,string"`
	Status       string
	HashRate     uint64 `json:"hashrate,string"`
	HashRateNice string `json:"hashrate_nice"`
}

type MyRigsResponse struct {
	Records []MyRigs `json:"records"`
}

func (c *Client) ListMyRigs() ([]MyRigs, error) {
	var data json.RawMessage
	params := getBasicMap("myrigs")
	_, err := c.Request("POST", "account", params, &data)
	if err != nil {
		return nil, err
	}

	var response MyRigsResponse
	if err := json.Unmarshal([]byte(data), &response); err != nil {
		return nil, err
	}
	return response.Records, err
}
