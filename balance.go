package miningrigrentals

import (
	"encoding/json"
)

type Balance struct {
	Confirmed   float64 `json:"confirmed,string"`
	Unconfirmed float64 `json:"unconfirmed,string"`
}

func (c *Client) GetBalance() (responses *Balance, err error) {
	var data json.RawMessage
	params := getBasicMap("balance")
	_, err = c.Request("POST", "account", params, &data)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(data), &responses); err != nil {
		return nil, err
	}
	return responses, err
}
