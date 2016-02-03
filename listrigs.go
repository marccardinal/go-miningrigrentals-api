package miningrigrentals

import (
	"encoding/json"
)

type RigListInfo struct {
	Start     uint32 `json:"start_num"`
	End       uint32 `json:"end_num"`
	Total     uint32 `json:"total,string"`
	Available RigListInfoStats `json:"available"`
	Rented    RigListInfoStats `json:"rented"`
	Price     RigListInfoPrice `json:"price"`
}

type RigListInfoStats struct {
	Rigs uint32 `json:"rigs,string"`
	Hash uint64 `json:"hash,string"`
}

type RigListInfoPrice struct {
	Lowest float64 `json:"lowest,string"`
	Last10 float64 `json:"last_10,string"`
	Last   float64 `json:"last,string"`
}

type RigList struct {
	Id           uint64 `json:"id,string"`
	Name         string `json:"name"`
	Rpi          string
	Type         string
	Online       uint8 `json:"online,string"`
	Price        float64 `json:"price,string"`
	PriceHour    float64 `json:"price_hr,string"`
	MinHours     uint16 `json:"minhrs,string"`
	MaxHours     uint16 `json:"maxhrs,string"`
	Rating       float64 `json:"rating,string"`
	Status       string
	HashRate     uint64 `json:"hashrate,string"`
	HashRateNice string `json:"hashrate_nice"`
}

type RigListResponse struct {
	Records []RigList `json:"records"`
	Info    RigListInfo `json:"info"`
}

func (c *Client) ListRigs(algotype string) ([]RigList, *RigListInfo, error) {
	var data json.RawMessage
	params := getBasicMap("list")
	params["type"] = algotype
	_, err := c.Request("POST", "rigs", params, &data)
	if err != nil {
		return nil, nil, err
	}
	//	fmt.Printf("%+v\n", string(data))

	var response RigListResponse
	if err := json.Unmarshal([]byte(data), &response); err != nil {
		return nil, nil, err
	}
	return response.Records, &response.Info, err
}

