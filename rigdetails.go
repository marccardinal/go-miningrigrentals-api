package miningrigrentals

import (
	"encoding/json"
	"strconv"
)


type RigDetails struct {
	Id               uint64 `json:"id,string"`
	Name             string `json:"name"`
	RentalId         uint64 `json:"rentalid,string,omitempty"`
	Rpi              string
	Owner            string
	StartTime        Time `json:"start_time,string"`
	Region           string
	Type             string
	Price            float64 `json:"price,string"`
	Status           string
	HashRate         RigHashRate `json:"hashrate"`
	MinHours         uint32
	MaxHours         uint32
	AvailableInHours float64 `json:"available_in_hours,string,omitempty"`
}

type rigDetails RigDetails

func (a *RigDetails) UnmarshalJSON(b []byte) error {
	var err error
	rd := rigDetails{}

	if err := json.Unmarshal(b, &rd); err != nil {
		return err
	}
	*a = RigDetails(rd)

	var f interface{}
	json.Unmarshal(b, &f)

	m := f.(map[string]interface{})

	hours := m["hours"]
	v := hours.(map[string]interface{})

	if a.MinHours, err = ParseUint16(v["min"].(string)); err != nil {
		return err
	}
	if a.MaxHours, err = ParseUint16(v["max"].(string)); err != nil {
		return err
	}

	return nil
}

type RigHashRate struct {
	Advertised  uint64 `json:"advertised,string"`
	HashRate5m  float64 `json:"5min"`
	HashRate30m float64 `json:"30min"`
	HashRate1h  float64 `json:"60min"`
}

func ParseUint16(s string) (uint16, error) {
	if v, err := strconv.ParseUint(s, 10, 16); err != nil {
		return 0, err
	} else {
		return uint16(v), nil
	}
}

func (c *Client) GetRigDetails(id int64) (*RigDetails, error) {
	var data json.RawMessage
	params := getBasicMap("detail")
	params["id"] = strconv.FormatInt(id, 10)
	_, err := c.Request("POST", "rigs", params, &data)
	if err != nil {
		return nil, err
	}

	var response *RigDetails
	if err := json.Unmarshal([]byte(data), &response); err != nil {
		return nil, err
	}
	return response, err
}


