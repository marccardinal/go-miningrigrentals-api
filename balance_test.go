package miningrigrentals

import (
	"reflect"
	"testing"
)

func Test_GetBalance_Success(t *testing.T) {
	server, client := testTools(200, `{"success":true,"version":"1","data":{"confirmed":"123.456","unconfirmed":"987.654"}}`)
	defer server.Close()
	responses, err := client.GetBalance()

	expect(t, err, nil)
	correctResponse := &Balance{
		Confirmed:123.456,
		Unconfirmed:987.654,
	}
	expect(t, reflect.DeepEqual(correctResponse, responses), true)
}