package miningrigrentals

import (
	"bytes"
	"crypto/hmac"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"crypto/sha1"
	"net/url"
	"encoding/hex"
	"log"
//	"net/http/httputil"
)

type Client struct {
	BaseURL    string
	Secret     string
	Key        string
	HTTPClient *http.Client
}

func New(key, secret string) *Client {
	client := Client{
		BaseURL:    "https://www.miningrigrentals.com/api/v1",
		Secret:     secret,
		Key:        key,
		HTTPClient: &http.Client{},
	}

	return &client
}

func GetNonce() int64 {
	now := time.Now()
	return now.UnixNano() / 1000
}

func GetNonceStr() string {
	nonce := GetNonce()
	return strconv.FormatInt(nonce, 10)
}

func getBasicMap(method string) map[string]string {
	answer := make(map[string]string)
	answer["method"] = method
	return answer
}

type Response struct {
	Success bool `json:"success"`
	Version int64 `json:"version,string"`
	Message string `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func (c *Client) Request(httpmethod string, apimethod string, params map[string]string, result *json.RawMessage) (res *http.Response, err error) {
	values := url.Values{}
	//turning input map into url.Values
	for k, v := range params {
		values.Set(k, v)
	}
	values.Set("nonce", GetNonceStr())

	sign, err := c.generateSig(values.Encode(), c.Secret);
	if err != nil {
		return res, err
	}

	fullURL := fmt.Sprintf("%s/%s", c.BaseURL, apimethod)
	req, err := http.NewRequest(httpmethod, fullURL, bytes.NewBufferString(values.Encode()))
	if err != nil {
		return res, err
	}

	//	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Go Library 0.1")
	req.Header.Add("x-api-key", c.Key)
	req.Header.Add("x-api-sign", sign)

	//	debug(httputil.DumpRequestOut(req, true))
	res, err = c.HTTPClient.Do(req)
	if err != nil {
		return res, err
	}
	defer res.Body.Close()
	//	debug(httputil.DumpResponse(res, true))

	if res.StatusCode != 200 {
		defer res.Body.Close()
		requestError := Error{}
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&requestError); err != nil {
			return res, err
		}
		return res, error(requestError)
	}

	if result != nil {
		var response Response
		decoder := json.NewDecoder(res.Body)
		if err = decoder.Decode(&response); err != nil {
			return res, err
		}
		if response.Success != true {
			return res, error(Error{Message:response.Message})
		}
		*result = response.Data
	}

	return res, nil
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}

func (c *Client) generateSig(message, secret string) (string, error) {
	/*
	key, err := hex.DecodeString(secret)
	if err != nil {
		return "", err
	}
	*/
	key := []byte(secret)

	signature := hmac.New(sha1.New, key)
	_, err := signature.Write([]byte(message))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(signature.Sum(nil)), nil
}