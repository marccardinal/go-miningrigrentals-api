package miningrigrentals

// from: https://github.com/preichenberger/go-coinbase-exchange/blob/master/time.go

import (
	"strings"
	"time"
)

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	var err error
	var parsedTime time.Time

	if string(data) == "null" {
		*t = Time(time.Time{})
		return nil
	}

	layouts := []string{
		"2006-01-02 15:04:05 UTC",
		"2006-01-02 15:04:05+00",
		"2006-01-02T15:04:05.999999Z",

		"2006-01-02 15:04:05.999999",
		"2006-01-02T15:04:05Z",
		"2006-01-02 15:04:05.999999+00"}
	for _, layout := range layouts {
		parsedTime, err = time.Parse(layout,
			strings.Replace(string(data), "\"", "", -1))
		if err != nil {
			continue
		}
		break
	}
	if parsedTime.IsZero() {
		return err
	}

	*t = Time(parsedTime)

	return nil
}

func (t *Time) Time() time.Time {
	return time.Time(*t)
}
