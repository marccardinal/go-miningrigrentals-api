package miningrigrentals

import (
	"encoding/json"
)

type MyRentals struct {
	Id               uint64 `json:"id,string"`
	RigId            uint64 `json:"rigid,string"`
	Name             string `json:"name"`
	StartTime        Time `json:"start_time,string"`
	Type             string
	Online           uint8 `json:"online,string"`
	Price            float64 `json:"price,string"`
	AvailableInHours float64 `json:"available_in_hours,string"`
	Status           string
	Advertised       MyRentalsHashRateAdvertised  `json:"advertised"`
	Current          MyRentalsHashRateCurrent  `json:"current"`
	Average          MyRentalsHashRateAverage `json:"average"`
}

type MyRentalsHashRateAdvertised struct {
	HashRate     float64 `json:"hashrate,string"`
	HashRateNice string `json:"hashrate_nice"`
}

type MyRentalsHashRateCurrent struct {
	HashRate5m      float64 `json:"hash_5m"`
	HashRate5mNice  string `json:"hash_5m_nice"`
	HashRate30m     float64 `json:"hash_30m"`
	HashRate30mNice string `json:"hash_30m_nice"`
	HashRate1h      float64 `json:"hash_1h"`
	HashRate1hNice  string `json:"hash_1h_nice"`
}

type MyRentalsHashRateAverage struct {
	HashRate     float64 `json:"hashrate"`
	HashRateNice string `json:"hashrate_nice"`
	Percent      float32 `json:"percent,string"`
}

type MyRentalsResponse struct {
	Records []MyRentals `json:"records"`
}

func (c *Client) ListMyRentals() ([]MyRentals, error) {
	var data json.RawMessage
	params := getBasicMap("myrentals")
	_, err := c.Request("POST", "account", params, &data)
	if err != nil {
		return nil, err
	}

	var response MyRentalsResponse
	if err := json.Unmarshal([]byte(data), &response); err != nil {
		return nil, err
	}
	return response.Records, err
}
