package miningrigrentals

import (
	"encoding/json"
	"strconv"
)

type RentalDetails struct {
	Id        uint64 `json:"id,string"`
	RigId     uint64 `json:"rigid,string"`
	Name      string `json:"name"`
	StartTime Time `json:"start_time,string"`
	Owner     string
	Renter    string
	Region    string
	Type      string
	HashRate  RentalHashrate `json:"hashrate"`
	Price     float64 `json:"price,string"`
	Length    uint64 `json:"length,string"`
	Left      float64 `json:"left,string"`
	Status    string
}

type RentalHashrate struct {
	Advertised  uint64 `json:"advertised,string"`
	Average     string
	HashRate5m  float64 `json:"5min"`
	HashRate30m float64 `json:"30min"`
	HashRate1h  float64 `json:"60min"`
}


func (c *Client) GetRentalDetails(id int64) (*RentalDetails, error) {
	var data json.RawMessage
	params := getBasicMap("detail")
	params["id"] = strconv.FormatInt(id, 10)
	_, err := c.Request("POST", "rental", params, &data)
	if err != nil {
		return nil, err
	}

	var response *RentalDetails
	if err := json.Unmarshal([]byte(data), &response); err != nil {
		return nil, err
	}
	return response, err
}

